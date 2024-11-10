import React, { useState, useEffect } from 'react';
import { FaHome, FaWifi, FaCoffee, FaSwimmingPool, FaParking } from 'react-icons/fa';
import { CgGym } from "react-icons/cg";
import { MdEdit } from 'react-icons/md';
import { FaPlus } from 'react-icons/fa';
import { Link } from 'react-router-dom';
import './home.css';

const MisHoteles = () => {
  const [hoteles, setHoteles] = useState([]);

  useEffect(() => {
    cargarHoteles(); 
  }, []);

  // cargue algunos hoteles a modo de prueba para ver como va a quedar 
  const cargarHoteles = () => {
    const data = [
      { id: 1, name: 'Hotel Boutique', description: 'Un lugar elegante y confortable en el corazón de la ciudad.', imageUrl: 'https://via.placeholder.com/150', amenities: ['WiFi', 'Cafe', 'Pileta'] },
      { id: 2, name: 'Hotel Lujo', description: 'Disfruta de una experiencia de lujo con vistas impresionantes.', imageUrl: 'https://via.placeholder.com/150', amenities: ['WiFi', 'Estacionamiento', 'Gimnasio'] },
      { id: 3, name: 'Hotel Económico', description: 'Una opción cómoda y accesible para viajeros con presupuesto limitado.', imageUrl: 'https://via.placeholder.com/150', amenities: ['WiFi', 'Cafe'] }
    ];
    setHoteles(data); 
  };

  return (
    <div className="contenedor-reserva">
      <h1>Reservar</h1>
      

      <div className="Barra-busqueda">
        <input type="text" placeholder="Busque su hotel aqui" />
        <div className="date-picker">
          <input className="date-field" type="date" />
          <input className="date-field" type="date" />
        </div>
        <button className="Agregar-Hotel">
          <FaPlus />
        </button>
      </div>
      <ul className="Grilla-amenities">
        {hoteles.length > 0 ? (
          hoteles.map((data) => (
            <li key={data.id} className="bloque">
              <img src={data.imageUrl} alt={data.name} className="hotel-imagen" />
              <h2>{data.name}</h2>
              <p>{data.description}</p>
              <div className="amenities">
                <h3>Amenidades:</h3>
                <ul className="Lista-amenities">
                  {data.amenities.map((amenity, index) => (
                    <li key={index} className="amenities">
                      {amenity === 'WiFi' && <FaWifi />}
                      {amenity === 'Cafe' && <FaCoffee />}
                      {amenity === 'Pileta' && <FaSwimmingPool />}
                      {amenity === 'Gimnasio' && <CgGym  />}
                      {amenity === 'Estacionamiento' && <FaParking />}
                      {` ${amenity}`}
                    </li>
                  ))}
                </ul>
              </div>
              <div className="boton-reserva">
                <button>Reservar</button>
                <button className="boton-editar">
                  <MdEdit />
                </button>
              </div>
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
