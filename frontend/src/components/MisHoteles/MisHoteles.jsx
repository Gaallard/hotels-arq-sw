import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './MisHoteles.css';
import { FaHome } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';
import { updateReserva, deleteReserva,tokenId, getreservas } from '../../utils/Acciones';

const MisHoteles = () => {
  const [hotels, setMyHotels] = useState([]);
  //const [reservaData, setInfo] = useState
  const [valorID,setID] = useState('');
  const [selectedHotel, setSelectedHotel] = useState(null);
  const [showEditDialog, setShowEditDialog] = useState(false);
  const [mensaje, setMensaje] = useState('');
  const [cantNoches, setCantNoches] = useState(1);
  const [reservas, setReservas] = useState('');


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
    if (!cantNoches || cantNoches <= 0) {
      setMensaje('Por favor, selecciona una cantidad válida de noches.');
      return;
    }
    if (selectedHotel) {
      const reservaData = {
        hotel_id: selectedHotel._id,
        user_id: await tokenId(),
        noches: cantNoches,
      };
      console.log("info recivida: ",reservaData)
      try {
        const newReserva = await updateReserva(reservaData);  // Llama a la función reserva pasando los datos
        setReservas((reservas) => [...reservas, newReserva]); 
        setMensaje('Reserva actualizada con éxito');
        closeEditDialog();
      } catch (error) {
        console.error('Error al actualizar la reserva:', error);
        setMensaje('Error al actualizar la reserva');
      }
    }
    window.location.reload();
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
              <h4>{data.noches}</h4> 
              <div className="boton-container">
                <button className="boton-actualizar" onClick={() => openEditDialog(data)}>
                  Actualizar
                </button>
                <button
            onClick={async () => {
              await deleteReserva(data.id);
              window.location.reload(); 
            }}
            className="boton-eliminar"
          >
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
              <input type="number" value={cantNoches} onChange={(e) => setCantNoches(Number(e.target.value))} placeholder="Edicar cantidad de noches" min={1}/>
              <div className="botones-mishoteles">
                <button onClick={closeEditDialog}>Cancelar</button>
                <button onClick={() => handleUpdateReserva(selectedHotel._id)}>Confirmar</button>
              </div>
            </div>
          </form>
        </div>
      )}  

    </div>
  ); 
};

export default MisHoteles;