import React, { useState, useEffect } from 'react';
import { FaHome, FaWifi, FaCoffee, FaSwimmingPool, FaParking } from 'react-icons/fa';
import { CgGym } from "react-icons/cg";
import { MdEdit } from 'react-icons/md';
import { FaPlus } from 'react-icons/fa';
import { Link } from 'react-router-dom';
import './homeAdmin.css';

const MisHoteles = () => {
  const [hoteles, setHoteles] = useState([]);

  useEffect(() => {
    cargarHoteles(); 
  }, []);

  // Carga de hoteles con datos adicionales
  const cargarHoteles = () => {
    const data = [
      { 
        id: 1, 
        name: 'Hotel Boutique', 
        description: 'Un lugar elegante y confortable en el corazón de la ciudad.', 
        imageUrl: 'https://via.placeholder.com/150', 
        amenities: ['WiFi', 'Cafe', 'Pileta'],
        address: '123 Main St', 
        city: 'Ciudad Principal', 
        state: 'Estado Central', 
        rating: 4.5, 
        price: '$150', 
        availableRooms: 5 
      },
      { 
        id: 2, 
        name: 'Hotel Lujo', 
        description: 'Disfruta de una experiencia de lujo con vistas impresionantes.', 
        imageUrl: 'https://via.placeholder.com/150', 
        amenities: ['WiFi', 'Estacionamiento', 'Gimnasio'],
        address: '456 High St', 
        city: 'Ciudad Lujosa', 
        state: 'Estado Norte', 
        rating: 5, 
        price: '$250', 
        availableRooms: 2 
      },
      { 
        id: 3, 
        name: 'Hotel Económico', 
        description: 'Una opción cómoda y accesible para viajeros con presupuesto limitado.', 
        imageUrl: 'https://via.placeholder.com/150', 
        amenities: ['WiFi', 'Cafe'],
        address: '789 Budget Rd', 
        city: 'Ciudad Económica', 
        state: 'Estado Sur', 
        rating: 3.8, 
        price: '$80', 
        availableRooms: 10 
      }
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
              <p><strong>Dirección:</strong> {data.address}, {data.city}, {data.state}</p>
              <p><strong>Calificación:</strong> {data.rating} / 5</p>
              <p><strong>Precio por noche:</strong> {data.price}</p>
              <p><strong>Habitaciones disponibles:</strong> {data.availableRooms}</p>
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