import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './MisHoteles.css';
import { FaHome } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';
import { updateReserva, deleteReserva,tokenId } from '../../utils/Acciones';

const MisHoteles = () => {
  const [hotels, setMyHotels] = useState([]);
  const [valorID,setID] = useState('');
  const val = async() =>{
      const val1 = await tokenId();
      setID(val1);
  }; 
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

  /*
  useEffect(() => {
    const token = sessionStorage.getItem('token');
    if (!token) {
      navigate('/users');
    } else {
      cargarHoteles(); // Cargar cursos inscritos si el usuario está autenticado
    }
  }, [navigate]);*/

  console.log("valor de val: ",valorID)
  useEffect(() => {
      fetch(`http://localhost:8083/reserva/misreservas/${valorID}`, {
          headers: {
              'Authorization': `Bearer ${localStorage.getItem('token')}`
          }
      })
      .then(response => response.json())
          .then(data => setMyHotels(data.results))
          .catch(error => {
              console.error('Error fetching courses:', error.message);
              console.error('Error details:', error.response);
          });
  },[valorID]);

  return (
    <div className="contenedor-misreservas">
      <h1>Mis Reservas</h1>
      <Link to='/home'>
        <button type="boton2" className="boton-casa">
          <FaHome className="icon" />
        </button>
      </Link>
      <ul className="grilla-hoteles">
        {hotels.length > 0 ? (
          hotels.map((data) => (
            <li key={data.id} className="lista-hoteles">
              <h2>{data.name}</h2>
              <p>{data.country}</p>
              <p>{data.state}</p>
              <p>{data.city}</p>
              <p>Precio por noche: {data.price}</p>
            
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
