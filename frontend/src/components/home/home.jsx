  /*
  const [hoteles, setHoteles] = useState([]);
  const navigate = useNavigate();
  const [isAdmin, setRole] = useState('');
  //const [busqueda, setBusqueda] = useState('');
  //const [cursosFiltrados, setCursosFiltrados] = useState([]);

  useEffect(() => {
    const token = sessionStorage.getItem('token');
    if (!token) {
      navigate('/users');
    } else {
      cargarHoteles(); // Cargar cursos si el usuario está autenticado
    }
  }, [navigate]);

  const cargarHoteles = async () => {
    try {
      const token = sessionStorage.getItem('token');
      const response = await fetch('http://localhost:8080/hotels', {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      if (!response.ok) {
        throw new Error('Error en la carga de cursos');
      }

      const data = await response.json();
      setHoteles(data); // Asignar los cursos obtenidos al estado cursos
      //setHotelsFiltrados(data); // Mostrar todos los cursos al inicio
    } catch (error) {
      console.error("Error durante la carga de hoteles:", error);
      alert('Error durante la carga de hoteles');
    }
  };

  const updateHotel = async () => {
    try {
      const token = sessionStorage.getItem('token');
      const response = await fetch('http://localhost:8080/hotels', {
        method: 'PUT',
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      if (!response.ok) {
        throw new Error('Error en la carga de cursos');
      }

      const data = await response.json();
      setHoteles(data); // Asignar los cursos obtenidos al estado cursos
      //setHotelsFiltrados(data); // Mostrar todos los cursos al inicio
    } catch (error) {
      console.error("Error durante la carga de cursos:", error);
      alert('Error durante la carga de cursos');
    }
  };

  const reserva = async (data) => {
    try {
      const token = sessionStorage.getItem('token');
      const idHotel = data.id; // Suponiendo que el objeto curso tiene un atributo id que representa el ID del curso
      const url = `http://localhost:8083/reservas/${idHotel}`;
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        }
      });
      if (response.status === 201) {
        Swal.fire('Inscripción exitosa', 'Bienvenido al curso', 'success');
      } else {
        Swal.fire('Error', 'No se pudo inscribir al curso', 'error');
      }
    } catch (error) {
      console.error("Error durante la inscripción:", error);
      alert('Error durante la inscripción');
    }
  };*/

  /*
  {isAdmin && (
              <Link to='/home' className='linkDelete'>
                <button onClick={() => DeleteCurso(courses.id)} className="delete">Eliminar curso</button>
              </Link>
            )}
               */

import React, { useState, useEffect } from 'react';
import { FaHome, FaWifi, FaCoffee, FaSwimmingPool, FaParking } from 'react-icons/fa';
import { insertHotel, updateHotel, reserva } from '../../utils/Acciones.js';
import { CgGym } from "react-icons/cg";
import { MdEdit } from 'react-icons/md';
import { FaPlus } from 'react-icons/fa';
import './home.css';
import { useNavigate } from 'react-router-dom';
import {tokenRole} from '../../utils/Acciones';
import {tokenId} from '../../utils/Acciones';


const MisHoteles = () => {
  const [showAddDialog, setShowAddDialog] = useState(false);
  const [showEditDialog, setShowEditDialog] = useState(false);

  const [selectedHotel, setSelectedHotel] = useState(null);

  // Abre el modal de agregar
  const openAddDialog = () => setShowAddDialog(true);
    
  // Cierra el modal de agregar
  const closeAddDialog = () => setShowAddDialog(false);

  // Abre el modal de editar con el hotel seleccionado
  const openEditDialog = (hotel) => {
    setSelectedHotel(hotel);
    setShowEditDialog(true);
  };

  // Cierra el modal de editar
  const closeEditDialog = () => {
    setSelectedHotel(null);
    setShowEditDialog(false);
  };

  const [name, setName] = useState('');
  const [address, setAddress] = useState(''); 
  const [country, setCountry] = useState(''); 
  const [city, setCity] = useState(''); 
  const [state, setState] = useState(''); 
  const [amenities, setAmenities] = useState([]); 
  const [rating, setRating] = useState(''); 
  const [price, setPrice] = useState(''); 
  const [available_rooms, setAvailableRooms] = useState('');
  const [reserva, setReservas] = useState('');
  const[mensaje, setMensaje] = useState('');
  const navigate = useNavigate();
  

    const [hotels, setHotels] = useState([]);
    const [isAdmin, setRole] = useState('');

    useEffect(() => {
        const fetchRole = async () => {
            try {
                const role = await tokenRole();
                setRole(role)
          
                console.log("role: ", role);
            } catch (error) {
                console.error('Error fetching role:', error);
            }
        };
        fetchRole();
    }, []);

    /*
    useEffect(() => {
      fetch(`http://localhost:8081/hotels`, {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
      })
        .then(response => {
          console.log("Raw response:", response);
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json();
        })
        .then(data => {
          setHotels(data.results); // Inicialmente muestra todos los hoteles
        })
        .catch(error => {
          console.error('Error fetching all hotels:', error.message);
        });
    }, []);*/

    useEffect(() => {
      fetch(`http://localhost:8081/hotels`, {
          headers: {
              'Authorization': `Bearer ${localStorage.getItem('token')}`
          }
      })
          .then(response => response.json())
          .then(data => setHotels(data.results))
          .catch(error => {
              console.error('Error fetching hotels:', error.message);
              console.error('Error details:', error.response);
          });
  }, []);

  const handleSubmitHotel = (e) => {
      e.preventDefault();
      const Data = { name, address, country, city, state, amenities, rating, price, available_rooms };

      insertHotel(Data).then(res => {
              setMensaje('Hotel creado exitosamente.');
              localStorage.setItem('hotel name: ', name);
              navigate("/home"); // Redirige a la página principal o a otra página después de crear el hotel
          }).catch(err => {
              setMensaje('Error al crear hotel');
              console.log(err);
          });
  };
  

     // Función para la actualización del hotel
     const handleUpdateHotel = async () => {
      const updatedData = { name, address, country, city, state, amenities, rating, price, available_rooms };
      try {
        const updatedHotel = await updateHotel(selectedHotel.id, updatedData);
        setHotels(hotels.map(hotel => hotel.id === selectedHotel.id ? updatedHotel : hotel));
        closeEditDialog();
      } catch (error) {
        console.error('Error al actualizar el hotel:', error);
      }
    };

    const handleReserva = async (hotelName) => {
      const reservaData = { hotel_name: hotelName, user_id: await tokenId() };  // Define los datos de la reserva
      try {
        const newReserva = await reserva(reservaData);  // Llama a la función `reserva` pasando los datos
        setReservas((reservas) => [...reservas, newReserva]);  // Actualiza el estado de reservas con la nueva reserva
        setMensaje('Reserva realizada con éxito');  // Muestra el mensaje de éxito
      } catch (error) {
        console.error('Error al realizar la reserva:', error);
        setMensaje('Error al realizar la reserva');  // Muestra el mensaje de error
      }
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



        {isAdmin && (
          <button className="Agregar-Hotel" onClick={openAddDialog}>
          <FaPlus />
        </button>
        )}
      </div>
      <ul className="Grilla-amenities">
        {hotels.length > 0 ? (
          hotels.map((data) => (
            <li key={data.id} className="bloque">
              <img src={data.imageUrl} alt={data.name} className="hotel-imagen" />
              <h2>{data.name}</h2>
              <p>{data.city}</p>
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
                <button onClick={() => reserva(data)}>Reservar</button>
              </div>

             {isAdmin &&(
              <button className="boton-editar" onClick={() => openEditDialog(data)}>
              <MdEdit />
            </button>
            ) } 
            </li>
          ))
        ) : (
          <p>No tienes hoteles disponibles</p>
        )}
      </ul>

      {showAddDialog && (
    <div className="modal">
      <form onSubmit={handleSubmitHotel}>
       <div className="modal-content">
        <h2>Agregar Nuevo Hotel</h2>
        <input type="text" id="name" value={name} onChange={(e) => setName(e.target.value)} placeholder="Nombre del Hotel" />
        <input type="text" id="Dirección" value={address} onChange={(e) => setAddress(e.target.value)} placeholder="Dirección" />
        <input type="text" id="Pais" value={country} onChange={(e) => setCountry(e.target.value)} placeholder="Pais" />
        <input type="text" id="Ciudad" value={city} onChange={(e) => setCity(e.target.value)} placeholder="Ciudad" />
        <input type="text" id="Estado" value={state} onChange={(e) => setState(e.target.value)} placeholder="Estado" />
        <input type="text" id="amenities" value={amenities} onChange={(e) => setAmenities(e.target.value.split(','))} placeholder="Amenities" />
        <input type="number" id="Calificación" value={rating} onChange={(e) => setRating(e.target.value)} placeholder="Calificación" />
        <input type="text" id="Precio" value={price} onChange={(e) => setPrice(e.target.value)} placeholder="Precio" />
        <input type="number" id="Habitaciones Disponibles" value={available_rooms} onChange={(e) => setAvailableRooms(e.target.value)} placeholder="Habitaciones Disponibles" />
        <button onClick={closeAddDialog}>Agregar</button>
        <button onClick={closeAddDialog}>Cancelar</button>
      </div> 
      </form>
    </div>
  )}

  

  
  {showEditDialog && selectedHotel && (
    <div className="modal">
      <form onSubmit={handleUpdateHotel}>
        <div className="modal-content">
        <h2>Editar Hotel</h2>
        <input type="text" value={name} onChange={(e) => setName(e.target.value)} placeholder="Nombre del Hotel" />
            <input type="text" value={address} onChange={(e) => setAddress(e.target.value)} placeholder="Dirección" />
            <input type="text" value={country} onChange={(e) => setCountry(e.target.value)} placeholder="País" />
            <input type="text" value={city} onChange={(e) => setCity(e.target.value)} placeholder="Ciudad" />
            <input type="text" value={state} onChange={(e) => setState(e.target.value)} placeholder="Estado" />
            <input type="text" value={amenities} onChange={(e) => setAmenities(e.target.value.split(','))} placeholder="Amenities" />
            <input type="number" value={rating} onChange={(e) => setRating(e.target.value)} placeholder="Calificación" />
            <input type="text" value={price} onChange={(e) => setPrice(e.target.value)} placeholder="Precio" />
            <input type="number" value={available_rooms} onChange={(e) => setAvailableRooms(e.target.value)} placeholder="Habitaciones Disponibles" />
        <button onClick={closeEditDialog}>Guardar Cambios</button>
        <button onClick={closeEditDialog}>Cancelar</button>
      </div>
      </form>
      
    </div>
  )}

    </div>
  ); 

  
};

export default MisHoteles;
