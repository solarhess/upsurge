package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gopdf/commerce"
    "github.com/jung-kurt/gofpdf"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    _ "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/zsais/go-gin-prometheus"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

const STOREFRONT="electronics"

var (
    occUrl string
    opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
        Name: "gopdf_events_processed",
        Help: "The total number of processed events",
    })
    pdfsGenerated = promauto.NewCounter(prometheus.CounterOpts{
        Name: "gopdf_pdfs_generated",
        Help: "The total number of processed events",
    })
)


type PdfRenderer struct {
    pdf *gofpdf.Fpdf
}

func (p PdfRenderer) Render(w http.ResponseWriter) error {
    return p.pdf.Output(w)
}

func (PdfRenderer) WriteContentType(w http.ResponseWriter) {
    header := w.Header()
    if val := header["Content-Type"]; len(val) == 0 {
        header["Content-Type"] = []string{"applicaton/pdf"}
    }
}


func main() {

    // Ensure the GATEWAY_URL environment variable is set
    // typically this comes in from a secret for a bound servicefactory instance
    occUrl = os.Getenv("GATEWAY_URL")
    if occUrl == "" {
        panic("Failed to set GATEWAY_URL environment variable")
    }

    // Enable Prometheus metrics on port :8081/metrics
    // Kyma likes metrics on 8081
    promhttp.Handler()
    http.Handle("/metrics", promhttp.Handler())
    go http.ListenAndServe(":8081", nil)


    // Initialize the real API handlers
    r := gin.Default()
    p := ginprometheus.NewPrometheus("gin")
    p.Use(r)

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.GET("/cardTypes", func(c *gin.Context) {
        cardTypesUrl := fmt.Sprintf("%s/%s/%s",occUrl,STOREFRONT,"cardtypes")
        resp, err := http.Get(cardTypesUrl)
        if err != nil {
            c.String(500,"Internal server error %s", err)
            return
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            c.String(500,"Internal server error %s", err)
            return
        }

        c.Data(200, "application/json", body)
    })

    r.GET("/swagger", func(c *gin.Context) {
        cardTypesUrl := fmt.Sprintf("%s/swagger.json",occUrl)
        resp, err := http.Get(cardTypesUrl)
        if err != nil {
            c.String(500,"Internal server error %s", err)
            return
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            c.String(500,"Internal server error %s", err)
            return
        }

        c.Data(200, "application/json", body)
    })

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "hello": "World!",
            "me":gin.H{
                "this":"is",
                "things": []string{"test","for","me",},
            },
        })
    })

    r.POST("/event/:eventType", func(c *gin.Context) {

        // record that we received an event
        opsProcessed.Inc()

    	eventType := c.Param("eventType")
        value, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            c.String(500, "Error %s", err)
        }

        log.Printf("Event Posted %s -- %s",eventType, string(value))

        switch eventType {
        case "order.created":
            orderEvent := commerce.UnmarshalOrderCreatedEvent(value)
            log.Printf("Handling order.created. order number %s", orderEvent.OrderCode)
            break
        default:
            log.Printf("Unrecognized Event received: %s", eventType)
            break
        }

        c.String(200, "Event Received OK")
    })

    r.GET("/pdf", func(c *gin.Context) {
        pdfsGenerated.Inc()

        pdfId := c.Query("pdfid")
        orderId := c.Query("orderid")
        pdf := gofpdf.New("P", "mm", "A4", "")
        pdf.AddPage()
        pdf.SetFont("Arial", "B", 16)
        pdf.Cell(40, 10, "Order id: " + orderId + " PDF id: " + pdfId)
        c.Render(200, &PdfRenderer{pdf:pdf})
    })

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
