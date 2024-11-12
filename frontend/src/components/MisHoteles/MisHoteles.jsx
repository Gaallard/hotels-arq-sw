import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './MisHoteles.css';
import { FaHome } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';
import { updateReserva, deleteReserva,tokenId, getreservas } from '../../utils/Acciones';

const MisHoteles = () => {
  const [hotels, setMyHotels] = useState([]);
  const [valorID,setID] = useState('');
  useEffect(() => {
    const obtenerTokenId = async () => {
      try {
        const val1 = await tokenId();
        setID(val1);
      } catch (error) {
        console.error('Error al obtener tokenId:', error);
      }
    };
    obtenerTokenId();
  }, []);
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

  useEffect(() => {
    const fetchHotels = async () => {
      try {
        const hotelsData = await getreservas();
        console.log("Hoteles cargados:", hotelsData); // Verifica la estructura de los datos
        setMyHotels(hotelsData.results || hotelsData); // Ajustwa según la estructura
      } catch (error) {
        console.error('Error fetching hotels:', error);
      }
    };
    fetchHotels();
  }, []);

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
              <h4>{data.noches}</h4> 
              <div className="boton-container">
                <button className="boton-actualizar">
                  Actualizar
                </button>
                <button className="boton-eliminar">
                  Eliminar
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

//onClick={() => handleUpdateReserva(data.id)}
//  onClick={() => handleDeleteReserva(data.id)}