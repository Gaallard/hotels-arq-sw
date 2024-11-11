import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './MisHoteles.css';
import { FaHome } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';
const MisHoteles = () => {
  const [hoteles, setHoteles] = useState([]);
  const navigate = useNavigate();

  /*
  useEffect(() => {
    cargarHoteles(); 
  }, []);


  const cargarHoteles = () => {
    const data = [
      { id: 1, name: 'Hotel Boutique', description: 'Un lugar elegante y confortable en el corazón de la ciudad.' },
      { id: 2, name: 'Hotel Lujo', description: 'Disfruta de una experiencia de lujo con vistas impresionantes.' },
      { id: 3, name: 'Hotel Económico', description: 'Una opción cómoda y accesible para viajeros con presupuesto limitado.' }
    ];
    setHoteles(data); 
  };*/

  useEffect(() => {
    const token = sessionStorage.getItem('token');
    if (!token) {
      navigate('/users');
    } else {
      cargarHoteles(); // Cargar cursos inscritos si el usuario está autenticado
    }
  }, [navigate]);

  const cargarHoteles = async () => {
    try {
      const token = sessionStorage.getItem('token'); // Suponiendo que el objeto usuario tiene un atributo id que representa el ID del usuario
      const url = `http://localhost:8080/hotels/${token}`;
      const response = await fetch(url, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });
  
      if (!response.ok) {
        throw new Error('Error en la carga de cursos inscritos');
      }
  
      const data = await response.json();
      setHoteles(data); 
    } catch (error) {
      console.error("Error durante la carga de cursos inscritos:", error);
      alert('Error durante la carga de cursos inscritos');
    }
  };

  return (
    <div className="contenedor-misreservas">
      <h1>Mis Reservas</h1>
      <Link to='/home'>
        <button type="boton2" className="boton-casa">
          <FaHome className="icon" />
        </button>
      </Link>
      <ul className="grilla-hoteles">
        {hoteles.length > 0 ? (
          hoteles.map((data) => (
            <li key={data.id} className="lista-hoteles">
              <h2>{data.name}</h2>
              <p>{data.country}</p>
              <p>{data.state}</p>
              <p>{data.city}</p>
              <p>Price per nigth: {data.price}</p>
            
            </li>
          ))
        ) : (
          <p>No tienes hoteles disponibles</p>
        )}
      </ul>
    </div>
  ); 
};

export default MisHoteles;
