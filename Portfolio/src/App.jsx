import { useState } from 'react'
import Proyectos from './component/proyectos.jsx'

function App() {

  return (
    <>
      <div>
        {/* Encabezado */}
        <header>
          <h1>Portfolio Santos</h1>
          <p>Hola! Soy Santos y este es mi Portfolio</p>
        </header>

        {/* Proyectos */}

        <main>
          <h2>Proyectos</h2>
          <Proyectos />
        </main>

        {/* Pie de Pagina */}

        <footer>
          <p>Contactos:</p>
          <ul>
            <li></li>
            <li></li>
          </ul>
        </footer>
      </div>
    </>
  )
}

export default App
