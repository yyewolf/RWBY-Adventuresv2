import io from 'socket.io-client'

const socket = io(process.env.VUE_APP_NOT_SECRET_CODE, { transports: ["websocket"] });

export default socket;