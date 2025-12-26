# 🏠Scraping de Departamentos en Cordoba

Scraping es una herramienta diseñada para ayudar a los usuarios a encontrar el momento ideal para alquilar un departamento. La aplicación realiza un seguimiento inteligente de los precios de publicaciones seleccionadas de Zonaprop, almacenando su historial y alertando sobre bajas de precio.

## Funcionalidades Principal
- **Scraping Personalizado:** Búsqueda de departamentos aplicando filtros (precio, ambientes, ubicación).
- **Dashboard de Monitoreo:** Visualización de los departamentos encontrados con opción de "Seguir".
- **Historial de Precios:** Registro cronológico de los cambios de valor de cada inmueble guardado.
- **Alertas de Tendencia:** Notificación visual cuando un departamento baja de precio respecto a su valor inicial.
- **Historial gráfico de precios:** Con Plotly.
- **Integración con Bot de WhatsApp:** Para alertas en tiempo real.
- **Análisis comparativo:** Precio vs. Promedio de Zona.

## Stack Tecnológico

### Backend
- **Lenguaje:** Python.

### Frontend
- **Framework:** Streamlit.
- **Visualización:** Plotly.

## Requeriments

* **Selenium:** Herramienta para abrir el browser y buscar los datos.
* **Webdriver-manager:** Se encarga de que `Selenium` tenga el driver correcto para el browser que se use.
* **Streamlit:** Framework para crear la interfaz de usuario.
* **Pandas:** Libreria para manejar tablas de datos.
* **Plotly:** Libreria para crear graficos.
