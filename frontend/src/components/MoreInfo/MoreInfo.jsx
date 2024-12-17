import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { getHotelById, tokenRole, updateHotel, reserva, tokenId } from '../../utils/Acciones.js';
import './MoreInfo.css';
import Swal from 'sweetalert2';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

const MoreInfo = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [hotel, setHotel] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [isAdmin, setRole] = useState('');
  const [cantNoches, setCantNoches] = useState(1);
  const [showEditDialog, setShowEditDialog] = useState(false);

  const [name, setName] = useState('');
  const [address, setAddress] = useState('');
  const [country, setCountry] = useState('');
  const [city, setCity] = useState('');
  const [state, setState] = useState('');
  const [amenities, setAmenities] = useState([]);
  const [rating, setRating] = useState('');
  const [price, setPrice] = useState('');
  const [available_rooms, setAvailableRooms] = useState('');
  const [mensaje, setMensaje] = useState('');
  const [reservas, setReservas] = useState('');
  const [reservaRealizada, setReservaRealizada] = useState(false);

  const [fechaIngreso, setFechaIngreso] = useState(new Date())
  const [fechaSalida, setFechaSalida] = useState(() => {
    const tomorrow = new Date();
    tomorrow.setDate(tomorrow.getDate() + 1);
    return tomorrow;
  });


  const calculateNights = (start, end) => {
    const diffTime = Math.abs(end - start);
    return Math.ceil(diffTime / (1000 * 60 * 60 * 24));
  };

  useEffect(() => {
    if (fechaIngreso && fechaSalida) {
      const nights = calculateNights(fechaIngreso, fechaSalida);
      setCantNoches(nights);
    }
  }, [fechaIngreso, fechaSalida]);


  useEffect(() => {
    const fetchHotelDetails = async () => {
      try {
        const hotelData = await getHotelById(id);
        setHotel(hotelData);
        const role = await tokenRole();
        setRole(role);
      } catch (err) {
        setError('Error al cargar los detalles del hotel.');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchHotelDetails();
  }, [id]);

  const openEditDialog = (hotel) => {
    setName(hotel.name || '');
    setAddress(hotel.address || '');
    setCountry(hotel.country || '');
    setCity(hotel.city || '');
    setState(hotel.state || '');
    setAmenities(hotel.amenities || []);
    setRating(hotel.rating || '');
    setPrice(hotel.price || '');
    setAvailableRooms(hotel.available_rooms || '');
    setShowEditDialog(true);
  };
  console.log("fecha ingreso: ",fechaIngreso)

  const closeEditDialog = () => {
    setShowEditDialog(false);
  };

  const handleUpdateHotelSubmit = async (e) => {
    e.preventDefault();

    const hotelData = {
      name,
      address,
      country,
      city,
      state,
      amenities,
      rating: parseFloat(rating),
      price: parseFloat(price),
      available_rooms: parseInt(available_rooms, 10),
    };

    try {
      await updateHotel(id, hotelData);
      setHotel(hotelData);
      setShowEditDialog(false);
    } catch (error) {
      console.error('Error al actualizar el hotel:', error);
    }
  };

  useEffect(() => {
    const reservaGuardada = localStorage.getItem(`reservaRealizada-${id}`);
    if (reservaGuardada === 'true') {
      setReservaRealizada(true);
    }
  }, [id]);
  
  const handleReserva = async (hotelId) => {
    if (!cantNoches || cantNoches <= 0) {
      setMensaje('Por favor, selecciona una cantidad válida de noches.');
      return;
    }
  
    const formatFecha = (fecha) => fecha.toISOString().slice(0, -5) + "Z";
  
    console.log("Fecha de ingreso:", fechaIngreso);
    console.log("Fecha de salida:", fechaSalida);

    const reservaData = {
      hotel_id: hotelId,
      noches: cantNoches,
      fecha_ingreso: fechaIngreso.toISOString().split(".")[0] + "Z", // Quita fracción de segundos
      fecha_salida: fechaSalida.toISOString().split(".")[0] + "Z",  // Quita fracción de segundos
      estado: 1,
    };

    console.log("mogolico:", reservaData.fecha_ingreso);
    console.log("mogolico 2:", reservaData.fecha_salida);
    
    try {
      console.log("Datos enviados para reserva:", reservaData);
      const newReserva = await reserva(reservaData);
      setReservas((prev) => [...prev, newReserva]);
      setMensaje('Reserva realizada con éxito');
      setReservaRealizada(true);
      localStorage.setItem(`reservaRealizada-${id}`, 'true');
      Swal.fire({
        icon: 'success',
        title: 'Reserva completada',
        text: '¡Su reserva se ha realizado con éxito!',
        confirmButtonText: 'Aceptar',
      });
    } catch (error) {
      console.error('Error al realizar la reserva:', error.response?.data || error.message);
      setMensaje('Error al realizar la reserva');
      Swal.fire({
        icon: 'error',
        title: 'Error en la reserva',
        text: 'Por favor, verifique las fechas e intente nuevamente.',
        confirmButtonText: 'Aceptar',
      });
    }
  };
  
  

  const handleFechaIngresoChange = (date) => {
    setFechaIngreso(date);
    const nextDay = new Date(date);
    nextDay.setDate(nextDay.getDate() + 1);
    if (fechaSalida <= date) {
      setFechaSalida(nextDay);
    }
  };
  


  if (loading) return <p>Cargando detalles del hotel...</p>;
  if (error) return <p>{error}</p>;

  return (
    <div className="hotel-details">
      <h1>{hotel.name}</h1>
      <p><strong>Dirección:</strong> {hotel.address}</p>
      <p><strong>Ciudad:</strong> {hotel.city}</p>
      <p><strong>País:</strong> {hotel.country}</p>
      <p><strong>Amenities:</strong> {hotel.amenities?.join(', ') || 'Ninguno'}</p>
      <p><strong>Calificación:</strong> {hotel.rating}</p>
      <p><strong>Precio por noche:</strong> {hotel.price}</p>
      <p><strong>Habitaciones disponibles:</strong> {hotel.available_rooms}</p>

      <div className="calendar-container">
        <label>Fecha de ingreso:</label>
        <DatePicker selected={fechaIngreso} onChange={handleFechaIngresoChange} minDate={new Date()}/>
        <label>Fecha de egreso:</label>
        <DatePicker selected={fechaSalida} onChange={(date) => setFechaSalida(date)} minDate={fechaIngreso} />
        <p><strong>Noches:</strong> {cantNoches}</p>
        <p><strong>Precio:</strong> {hotel.price * cantNoches}</p>
      </div>

      <div className="botones_reserva">       
        <button onClick={() => navigate(-1)}>Volver</button>
        {!reservaRealizada && (
          <button onClick={() => handleReserva(hotel.id)}>Reservar</button>
        )}
        {isAdmin && (
          <button onClick={() => openEditDialog(hotel)}>Editar</button>
        )}
      </div>

      {showEditDialog && (
        <div className="modal">
          <form onSubmit={handleUpdateHotelSubmit}>
            <div className="modal-content">
              <h2>Editar Hotel</h2>
              <input type="text" value={name} onChange={(e) => setName(e.target.value)} placeholder="Nombre del Hotel" />
              <input type="text" value={address} onChange={(e) => setAddress(e.target.value)} placeholder="Dirección" />
              <input type="text" value={country} onChange={(e) => setCountry(e.target.value)} placeholder="País" />
              <input type="text" value={city} onChange={(e) => setCity(e.target.value)} placeholder="Ciudad" />
              <input type="text" value={state} onChange={(e) => setState(e.target.value)} placeholder="Estado" />
              <input type="text" value={amenities} onChange={(e) => setAmenities(e.target.value.split(','))} placeholder="Amenities" />
              <input type="number" value={rating} onChange={(e) => setRating(e.target.value)} placeholder="Calificación" />
              <input type="text" value={price} onChange={(e) => setPrice(e.target.value)} placeholder="Precio por noche" />
              <input type="number" value={available_rooms} onChange={(e) => setAvailableRooms(e.target.value)} placeholder="Habitaciones Disponibles" />
              <button type="button" onClick={closeEditDialog}>Cancelar</button>
              <button type="submit">Confirmar</button>
            </div>
          </form>
        </div>
      )}
    </div>
  );
};

export default MoreInfo;