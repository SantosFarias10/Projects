package scraper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"

	"github.com/SantosFarias10/deptoHunter/service"
)

func ScrapearSitio(url string) {
	fmt.Printf("Iniciando scraper en: %s\n", url)

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second,
	})

	c.OnError(func(res *colly.Response, err error) {
		fmt.Printf("Error %d en %s: %s\n", res.StatusCode, res.Request.URL, err)
	})

	c.OnScraped(func(res *colly.Response) {
		fmt.Printf("Scraping Finalizado: %s\n", res.Request.URL)
	})

	if strings.Contains(url, "argenprop.com") {
		configuracionArgenprop(c)
	} else if strings.Contains(url, "remax.com.ar") {
		configuracionRemax(c)
	} else if strings.Contains(url, "zonaprop.com.ar") {
		configuracionZonaProp(c)
	} else {
		fmt.Printf("Dominio no soportado: %s\n", url)
	}

	c.Visit(url)
}

func configuracionArgenprop(c *colly.Collector) {
	fmt.Println("Usando mapa de navegación: ArgenProp (Modo Iconos)")

	c.OnHTML("div.listing__item", func(e *colly.HTMLElement) {
		link := "https://www.argenprop.com" + e.ChildAttr("a", "href")

		titulo := strings.TrimSpace(e.ChildText(".card__title"))
		if titulo == "" {
			titulo = strings.TrimSpace(e.ChildText(".card__title--primary"))
		}

		ubicacion := strings.TrimSpace(e.ChildText(".card__address"))
		precioRaw := e.ChildText(".card__price")

		var habitaciones int
		var baños int
		var metros float64

		e.ForEach("ul.card__main-features li, ul.card__common-data li", func(_ int, el *colly.HTMLElement) {

			iconoClass := el.ChildAttr("i", "class")
			texto := strings.ToLower(el.Text)

			if strings.Contains(iconoClass, "dormitorios") {
				habitaciones = convertirNumero(texto)
			} else if strings.Contains(iconoClass, "banos") || strings.Contains(iconoClass, "baños") {
				baños = convertirNumero(texto)
			} else if strings.Contains(iconoClass, "superficie") {
				metros = convertirFloat(texto)
			} else {
				if strings.Contains(texto, "dorm") || strings.Contains(texto, "amb") {
					habitaciones = convertirNumero(texto)
				} else if strings.Contains(texto, "baño") || strings.Contains(texto, "toile") {
					baños = convertirNumero(texto)
				} else if strings.Contains(texto, "m²") || strings.Contains(texto, "sup") || strings.Contains(texto, "mts") {
					metros = convertirFloat(texto)
				}
			}
		})

		tituloLower := strings.ToLower(titulo)

		if habitaciones == 0 {
			if strings.Contains(tituloLower, "monoamb") || strings.Contains(tituloLower, "ambiente unico") {
				habitaciones = 1
			} else {
				habitaciones = extraerNumeroConRegex(tituloLower, `(\d+)\s*(dorm|amb)`)
			}
		}

		if baños == 0 {
			baños = extraerNumeroConRegex(tituloLower, `(\d+)\s*(baño|bano)`)
			if baños == 0 && habitaciones > 0 {
				baños = 1
			}
		}

		if metros == 0 {
			metros = extraerFloatConRegex(tituloLower, `(\d+[.,]?\d*)\s*(m2|m²|mts|mt|sup|cub|tot|metros)`)
		}

		precio, moneda := extraerPrecio(precioRaw)

		if link != "https://www.argenprop.com" {
			service.ProcesarDepto(titulo, link, ubicacion, habitaciones, baños, metros, precio, moneda)
		}
	})
}

func configuracionRemax(c *colly.Collector) {
	fmt.Println("Usando mapa de navegación: RE/MAX")
	c.OnHTML("div.gallery-item-container", func(e *colly.HTMLElement) {
		link := "https://www.remax.com.ar" + e.ChildAttr("a", "href")
		titulo := e.ChildText("h3.card-title")
		precioRaw := e.ChildText(".price")
		ubicacion := e.ChildText(".card-location")

		var habitaciones int
		var baños int
		var metros float64

		e.ForEach(".card-feature, .feature-item", func(_ int, el *colly.HTMLElement) {
			texto := strings.ToLower(el.Text)
			if strings.Contains(texto, "dorm") || strings.Contains(texto, "amb") {
				habitaciones = convertirNumero(texto)
			} else if strings.Contains(texto, "baño") {
				baños = convertirNumero(texto)
			} else if strings.Contains(texto, "m²") || strings.Contains(texto, "sup") {
				metros = convertirFloat(texto)
			}
		})
		precio, moneda := extraerPrecio(precioRaw)
		service.ProcesarDepto(titulo, link, ubicacion, habitaciones, baños, metros, precio, moneda)
	})
}

func configuracionZonaProp(c *colly.Collector) {
	fmt.Println("Usando mapa de navegación: ZonaProp")
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0")
	})
	c.OnHTML("div.postingCard", func(e *colly.HTMLElement) {
		rawLink := e.ChildAttr("a", "href")
		link := "https://www.zonaprop.com.ar" + rawLink
		if strings.Contains(rawLink, "http") {
			link = rawLink
		}

		titulo := e.ChildText("h2.postingCardTitle")
		precioRaw := e.ChildText(".postingCardPrice")
		if precioRaw == "" {
			precioRaw = e.ChildText(".firstPrice")
		}
		ubicacion := e.ChildText(".postingCardLocationTitle")

		var habitaciones int
		var baños int
		var metros float64

		e.ForEach("ul.postingCardMainFeatures li", func(_ int, el *colly.HTMLElement) {
			texto := strings.ToLower(el.Text)
			if strings.Contains(texto, "dorm") {
				habitaciones = convertirNumero(texto)
			} else if strings.Contains(texto, "baño") {
				baños = convertirNumero(texto)
			} else if strings.Contains(texto, "m²") {
				metros = convertirFloat(texto)
			}
		})
		precio, moneda := extraerPrecio(precioRaw)
		service.ProcesarDepto(titulo, link, ubicacion, habitaciones, baños, metros, precio, moneda)
	})
}

func limpiarTexto(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func convertirNumero(texto string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(texto)
	num, _ := strconv.Atoi(match)
	return num
}

func convertirFloat(texto string) float64 {
	re := regexp.MustCompile(`[0-9.,]+`)
	match := re.FindString(texto)

	if match == "" {
		return 0.0
	}

	match = strings.ReplaceAll(match, ".", "")
	match = strings.ReplaceAll(match, ",", ".")

	num, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0.0
	}
	return num
}

func extraerPrecio(precioRaw string) (float64, string) {
	precioRaw = limpiarTexto(precioRaw)
	moneda := "ARS"
	if strings.Contains(precioRaw, "USD") || strings.Contains(precioRaw, "U$S") {
		moneda = "USD"
	}
	precio := convertirFloat(precioRaw)
	return precio, moneda
}

func extraerNumeroConRegex(texto string, patron string) int {
	re := regexp.MustCompile(patron)
	match := re.FindStringSubmatch(texto)
	if len(match) > 1 {
		num, _ := strconv.Atoi(match[1])
		return num
	}
	return 0
}

func extraerFloatConRegex(texto string, patron string) float64 {
	re := regexp.MustCompile(patron)
	match := re.FindStringSubmatch(texto)
	if len(match) > 1 {
		return convertirFloat(match[1])
	}
	return 0.0
}
