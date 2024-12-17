import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './MisHoteles.css';
import { FaHome } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import { updateReserva, deleteReserva, tokenId, getreservas } from '../../utils/Acciones';

const MisHoteles = () => {
  const [hotels, setMyHotels] = useState([]);
  //const [reservaData, setInfo] = useState
  const [valorID, setID] = useState('');
  const [selectedHotel, setSelectedHotel] = useState(null);
  const [showEditDialog, setShowEditDialog] = useState(false);
  const [mensaje, setMensaje] = useState('');
  const [cantNoches, setCantNoches] = useState(1);
  const [reservas, setReservas] = useState('');
  const [fechaIngreso, setFechaIngreso] = useState(new Date())
  const [fechaSalida, setFechaSalida] = useState(new Date())


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

  const openEditDialog = (hotel) => {
    setSelectedHotel(hotel);
    setCantNoches(hotel.noches); // Establece el valor inicial de noches en el modal
    setShowEditDialog(true);
  };

  const closeEditDialog = () => {
    setSelectedHotel(null);
    setShowEditDialog(false);
  };

  const handleUpdateReserva = async (e) => {
    e.preventDefault();
    const noches = calculateNights(fechaIngreso, fechaSalida);

    if (!noches || noches <= 0) {
      setMensaje('Por favor, selecciona fechas válidas.');
      return;
    }

    if (selectedHotel) {
      const reservaData = {
        hotel_id: selectedHotel.id,
        user_id: await tokenId(),
        fechaIngreso: fechaIngreso.toISOString(),
        fechaSalida: fechaSalida.toISOString(),
        noches,
      };

      try {
        const newReserva = await updateReserva(reservaData);
        setReservas((reservas) => [...reservas, newReserva]);
        setMensaje('Reserva actualizada con éxito');
        closeEditDialog();
        window.location.reload();
      } catch (error) {
        console.error('Error al actualizar la reserva:', error);
        setMensaje('Error al actualizar la reserva');
      }
    }
  };

  return (
    <div className="contenedor-misreservas">
      <h1>MIS RESERVAS</h1>
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
              <h4>Fecha ingreso:  {data.fecha_ingreso}</h4> 
              <h4>Fecha egreso:  {data.fecha_salida}</h4> 
              <h4>Noches reservadas: {data.noches}</h4> 
              <h4>Precio de la reserva ${data.price * data.noches}</h4>
              <div className="boton-container">
                <button className="boton-actualizar" onClick={() => openEditDialog(data)}>
                  Actualizar
                </button>
                <button
            onClick={async () => {
              await deleteReserva(data.id);
              window.location.reload(); 
            }}
            className="boton-eliminar">
            Eliminar
          </button>
              </div>           
            </li>
          ))
        ) : (
          <p>No tienes hoteles disponibles</p>
        )}
      </ul>

      {showEditDialog && selectedHotel && (
        <div className="modal">
          <form onSubmit={handleUpdateReserva}>
            <div className="modal-content">
              <h2>Editar Reserva</h2>
              <label>Fecha de Ingreso:</label>
              <DatePicker
                selected={fechaIngreso}
                onChange={(date) => setFechaIngreso(date)}
                dateFormat="dd-MM-yyyy"
              />
              <label>Fecha de Salida:</label>
              <DatePicker
                selected={fechaSalida}
                onChange={(date) => setFechaSalida(date)}
                dateFormat="dd-MM-yyyy"
                minDate={fechaIngreso} // Evita seleccionar fechas anteriores a la de ingreso
              />
              <div className="botones-mishoteles">
                <button type="button" onClick={closeEditDialog}>Cancelar</button>
                <button type="submit">Confirmar</button>
              </div>
            </div>
          </form>
        </div>
      )}  

    </div>
  ); 
};

export default MisHoteles;