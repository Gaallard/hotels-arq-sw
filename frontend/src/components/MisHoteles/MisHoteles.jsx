import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './MisHoteles.css';
import { FaHome } from 'react-icons/fa';

const MisHoteles = () => {
  const [hoteles, setHoteles] = useState([]);

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
              <p>{data.description}</p>
            
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
