import { projects } from "../projects.jsx"

export default function Proyectos() {
  return (
    <div>
      {projects.map(project => (
        <div key={project.id}>
          <h3>{project.title}</h3>
          <p>{project.description}</p>
          <p>{project.tecnologies.join(', ')}</p>
          <img src={project.image} alt={project.title} />
          <br></br>
          <a href={project.link}>Ver proyecto</a>
        </div>
      ))}
    </div>
  )
}