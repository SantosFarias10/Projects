# Depto Hunter

## Descripción del Proyecto

Esta aplicación es una herramienta automatizada de inteligencia inmobiliaria diseñada para rastrear, almacenar y analizar la fluctuación de precios de departamentos en alquiler o venta.

A diferencia de los portales inmobiliarios tradicionales que muestran el precio actual, **este sistema construye un historial**, permitiendo al usuario identificar:
1. Bajadas de precio reales (oportunidades).
2. "Falsas ofertas" (precios inflados y luego rebajados).
3. Tendencias del mercado en zonas específicas.

El objetivo es empoderar al usuario con datos históricos para tomar la mejor decisión de alquiler o compra.

---

## Características Principales

* **Scraping Concurrente (Go):** Extracción masiva y rápida de datos utilizando Goroutines.
* **Historial de Precios:** Base de datos relacional que registra cada cambio de precio.
* **Notificaciones de Escritorio (Web Push):** El sistema envía una alerta nativa a tu computadora (PC/Mac) cuando detecta una bajada de precio, sin necesidad de tener la web abierta en primer plano.
* **Análisis de Valor ($/m^2$):** Cálculo automático del precio por metro cuadrado para identificar oportunidades reales ocultas tras precios totales altos.
* **Detección de Falsas Ofertas:** Algoritmo que ignora descuentos artificiales basados en subidas previas recientes.

...

## Futuras Mejoras
- [ ] Integración con mapas interactivos para ver precios por zonas (Heatmap).
- [ ] Sistema de login para guardar favoritos.

---

## Stack Tecnológico

Este proyecto utiliza una arquitectura moderna de cliente-servidor (Frontend y Backend separados).

### 🐹 Backend & Data (Go)
* **Lenguaje:** Go (Golang)
* **Scraping Engine:** `Playwright-Go` (Para renderizado de JS) o `Colly` (Para velocidad pura).
* **API Framework:** `Gin Gonic` (High-performance HTTP web framework).
* **Base de Datos:** `SQLite` (Local) con `GORM` (Object Relational Mapper).
* **Concurrencia:** Uso de Goroutines para scraping paralelo masivo.

### ⚛️ Frontend / UI (JavaScript)
* **Framework:** `React.js` (Vite para el entorno de desarrollo).
* **Estilos:** `Tailwind CSS` (Para un diseño rápido, responsivo y moderno).
* **Gráficos:** `Recharts` o `Chart.js` (Para visualizar la línea de tiempo de los precios).
* **Conexión:** `Axios` o `Fetch` (Para consumir la API de Python).

---

## 📂 Estructura del Proyecto (Tentativa)

```text
/real-estate-tracker
│
├── /backend
│   ├── /app
│   │   ├── /scrapers      # Lógica de extracción de datos
│   │   ├── /models        # Modelos de Base de Datos (SQLAlchemy)
│   │   ├── /routers       # Endpoints de la API (FastAPI)
│   │   └── main.py        # Punto de entrada del servidor
│   ├── requirements.txt
│   └── database.db
│
├── /frontend
│   ├── /src
│   │   ├── /components    # Tarjetas de propiedades, Gráficos
│   │   ├── /pages         # Dashboard, Detalle de propiedad
│   │   └── App.jsx
│   └── package.json
│
└── README.md
