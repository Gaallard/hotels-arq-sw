import React, { useState, useEffect } from "react";
import axios from "axios";

const App = () => {
  const [containers, setContainers] = useState([]);
  const token = localStorage.getItem('token');

  // Fetch de datos
  const fetchContainers = async () => {
    try {
      const response = await axios.get("http://localhost:8080/users/containers", {
        headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
      });
      setContainers(response.data);
    } catch (error) {
      console.error("Error al obtener los contenedores:", error);
    }
  };

  // Manejo de contenedores (start/stop)
  const handleAction = async (action, containerName) => {
    try {
      await axios.post(`http://localhost:8080/users/containers/${action}/${containerName}`, {
        headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
      });
      fetchContainers(); // Actualizar la lista
    } catch (error) {
      console.error(`Error al ${action} el contenedor ${containerName}:`, error);
    }
  };

  useEffect(() => {
    fetchContainers();
  }, []);

  return (
    <div>
      <h1>Monitor de Microservicios</h1>
      <table>
        <thead>
          <tr>
            <th>Nombre</th>
            <th>Estado</th>
            <th>Acciones</th>
          </tr>
        </thead>
        <tbody>
          {containers.map((container) => (
            <tr key={container.name}>
              <td>{container.name}</td>
              <td>{container.status}</td>
              <td>
                {container.status === "running" ? (
                  <button onClick={() => handleAction("stop", container.name)}>Detener</button>
                ) : (
                  <button onClick={() => handleAction("start", container.name)}>Iniciar</button>
                )}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default App;
