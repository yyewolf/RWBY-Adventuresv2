import io from 'socket.io-client'

const socket = io("ws://localhost:9003/", { transports: ["websocket"] });

export default socket;