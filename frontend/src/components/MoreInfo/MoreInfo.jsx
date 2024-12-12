

import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { getHotelById, tokenRole, register } from '../../utils/Acciones.js';
import './MoreInfo.css';

const MoreInfo = () => {
  const { id } = useParams(); // Obtiene el ID del hotel desde la URL
  const navigate = useNavigate();
  const [hotel, setHotel] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchHotelDetails = async () => {
      try {
        const hotelData = await getHotelById(id);
        console.log('id boliviano: ', id) // Llama a una función para obtener los detalles del hotel
        setHotel(hotelData);
      } catch (err) {
        setError('Error al cargar los detalles del hotel.');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchHotelDetails();
  }, [id]);

  if (loading) return <p>Cargando detalles del hotel...</p>;
  if (error) return <p>{error}</p>;

  return (
    <div className="hotel-details">
      <h1>{hotel.name}</h1>
      <p>{hotel.description || 'No hay descripción disponible.'}</p>
      <p><strong>Dirección:</strong> {hotel.address}</p>
      <p><strong>Ciudad:</strong> {hotel.city}</p>
      <p><strong>País:</strong> {hotel.country}</p>
      <p><strong>Amenities:</strong> {hotel.amenities?.join(', ') || 'Ninguno'}</p>
      <p><strong>Calificación:</strong> {hotel.rating}</p>
      <p><strong>Precio:</strong> {hotel.price}</p>
      <p><strong>Habitaciones disponibles:</strong> {hotel.available_rooms}</p>

      <button onClick={() => navigate(-1)}>Volver</button>
    </div>
  );
};

export default MoreInfo;


/*
import React, { useEffect, useState } from 'react';
import { getHotelById, reserva, tokenRole } from '../../utils/Acciones';
import { useParams } from 'react-router-dom'; // Supongo que usas react-router para navegación

const HotelDetails = () => {
  const { hotelId } = useParams(); // Obtenemos el ID del hotel desde la URL
  const [hotel, setHotel] = useState(null); // Estado para almacenar los datos del hotel
  const [isAdmin, setRole] = useState(false); // Estado para verificar si es admin
  const [nights, setNights] = useState(1); // Estado para manejar las noches de reserva
  const [status, setStatus] = useState("Confirmada"); // Estado para el estado de la reserva

  useEffect(() => {
    // Cargar detalles del hotel
    async function fetchHotelDetails() {
      try {
        const hotelData = await getHotelById(hotelId);
        setHotel(hotelData);
        console.log('hotel boliviano: ', hotelData)
      } catch (error) {
        console.error('Error al cargar los detalles del hotel:', error);
      }
    }

    // Verificar rol del usuario
    async function checkRole() {
      try {
        const role = await tokenRole();
        setRole(role);
      } catch (error) {
        console.error('Error al verificar el rol del usuario:', error);
      }
    }

    fetchHotelDetails();
    checkRole();
  }, [hotelId]);

  const handleReservation = async () => {
    try {
      await reserva({
        hotel_id: hotelId,
        noches: nights,
        estado: status,
      });
      alert('Reserva realizada exitosamente');
    } catch (error) {
      console.error('Error al realizar la reserva:', error);
      alert('Error al realizar la reserva');
    }
  };

  if (!hotel) {
    return <p>Cargando detalles del hotel...</p>;
  }

  return (
    <div>
      <h1>Detalles del Hotel</h1>
      <h2>{hotel.name}</h2>
      <p><strong>Dirección:</strong> {hotel.address}</p>
      <p><strong>Ciudad:</strong> {hotel.city}</p>
      <p><strong>Estado:</strong> {hotel.state}</p>
      <p><strong>País:</strong> {hotel.country}</p>
      <p><strong>Rating:</strong> {hotel.rating}</p>
      <p><strong>Amenidades:</strong> {hotel.amenities.join(', ')}</p>
      <p><strong>Precio por noche:</strong> ${hotel.price}</p>
      <p><strong>Habitaciones disponibles:</strong> {hotel.available_rooms}</p>

      {!isAdmin && (
        <div>
          <h3>Realizar Reserva</h3>
          <label>
            Noches:
            <input
              type="number"
              value={nights}
              min="1"
              onChange={(e) => setNights(Number(e.target.value))}
            />
          </label>
          <button onClick={handleReservation}>Reservar</button>
        </div>
      )}
    </div>
  );
};

export default HotelDetails;


/*import React, { useEffect, useState } from 'react';
import { buscarSuscription, DeleteCurso, suscribe, Docomment, tokenRole, uploadFile } from '../Utils/Acciones';
import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser, faHome } from '@fortawesome/free-solid-svg-icons';
import '../../src/Css/moreinfo.css';
import '../../src/Css/App.css';
import '../../src/Css/Account.css';

const CourseDetail = () => {
  const [courses, setCourse] = useState([]);
  const [isSubscribed, setIsSubscribed] = useState(false);
  const [comment, setComment] = useState('');
  const [Comments, getComments] = useState([]);
  const [isAdmin, setRole] = useState('');
  const [file, setFile] = useState();
  const [message, setMessage] = useState('');

  const getId = parseInt(localStorage.getItem('hotelid'));
  console.log('id del hotel: ', getId);

  useEffect(() => {
    const fetchRole = async () => {
      try {
        const role = await tokenRole();
        setRole(role);
        console.log('role: ', role);
      } catch (error) {
        console.error('Error fetching role:', error);
      }
    };
    fetchRole();
  }, []);

  useEffect(() => {
    const fetchCourse = async () => {
      try {
        const response = await fetch(`http://localhost:8080/courses/courseInfo/${getId}`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        });
        const data = await response.json();
        setCourse(data);
        console.log('hoola', data);

        const isUserSubscribed = await buscarSuscription(data.id);
        setIsSubscribed(!!isUserSubscribed);
      } catch (error) {
        console.error('Error fetching course:', error.message);
        console.error('Error details:', error.response);
      }
    };

    fetchCourse();
  }, [getId]);

  const handleSubscribe = async (courseId) => {
    try {
      await suscribe(courseId);
      setIsSubscribed(true);
      localStorage.setItem(`isSubscribed_${courseId}`, 'true');
    } catch (error) {
      console.error('Error al suscribirse al curso:', error);
      console.log('Error details:', error.response.data);
    }
  };

  useEffect(() => {
    fetchComments();
  }, [getId]);

  const fetchComments = async () => {
    try {
      const response = await fetch(`http://localhost:8080/comments/${getId}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
      });
      const data = await response.json();
      getComments(data);
    } catch (error) {
      console.error('Error al obtener comments:', error);
    }
  };

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleFileUpload = async (e) => {
    e.preventDefault();
    try {
      const response = await uploadFile(file, getId);
      setMessage(response.message || 'Archivo subido correctamente!');
    } catch (error) {
      setMessage('Error al subir el archivo');
    }
  };

  const handleSuscription = async (courseId) => {
    try {
      const suscp = await buscarSuscription(courseId);
      console.log('anduvo bien: ', suscp);
      const data = {
        suscription_id: suscp,
        comment: comment,
      };
      await Docomment(data);
      fetchComments();
      setComment('');
    } catch (error) {
      console.error('Error commenting:', error);
    }
  };

  console.log('role: ', isAdmin);

  return (
    <div className='courseMoreInfo'>
      <header className="navBar">
        <Link to="/home" className='linkPageTitle'>
          <h1 className='pageTitle'>Details</h1>
        </Link>
        <div className='Perfil-Home-Box'>
          <Link to='/myAccount' className='linkMyAccountButton'>
            <button className='myAccountButton'>
              <FontAwesomeIcon icon={faUser} /> Perfil
            </button>
          </Link>
          <Link to='/home' className='linkHomeButton'>
            <button className="buttonToHome">
              <FontAwesomeIcon icon={faHome} /> Home
            </button>
          </Link>
        </div>
      </header>
      <div className="info">
        <div key={courses.id} className="Course">
          <div className="CourseDetail">
            <h1 className='CourseTitle'>{courses.name}</h1>
            <p className='CourseDescription'>{courses.descriptionlarga}</p>
            <p className='CourseCategory'>Cateoria: {courses.category}</p>
            <p className='CourseCategory'>Profesor: {courses.profesor}</p>
            <p className='CourseCategory'>Horas de cursado: {courses.horas}</p>
            <button onClick={() => handleSubscribe(courses.id)} className="suscribe">
              {isSubscribed ? 'Suscripto' : 'Suscribirse'}
            </button>

            {isSubscribed && (
              <div className='Comentar'>
                <input
                  className='comentariotexto'
                  type="text"
                  placeholder='Comentario'
                  value={comment}
                  onChange={(e) => setComment(e.target.value)}
                />
                <button onClick={() => handleSuscription(courses.id)} className="suscribe">Comentar</button>
              </div>
            )}

            {isAdmin && (
              <Link to='/home' className='linkDelete'>
                <button onClick={() => DeleteCurso(courses.id)} className="delete">Eliminar curso</button>
              </Link>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

export default CourseDetail; */