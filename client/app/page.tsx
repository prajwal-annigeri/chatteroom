"use client";

import { useState, useEffect, useContext } from "react";
import { API_URL, WEBSOCKET_URL } from "../constants";
import { v4 as uuidv4 } from "uuid";
import { AuthContext } from "../modules/auth_provider";
import { WebSocketContext } from "../modules/websocket_provider";
import { useRouter } from "next/navigation";

const page = () => {
  const [rooms, setRooms] = useState<{ id: string; name: string }[]>([]);
  const [roomName, setRoomName] = useState("");
  const { user } = useContext(AuthContext);
  const { setConn } = useContext(WebSocketContext);

  const router = useRouter();

  const getRooms = async () => {
    try {
      const res = await fetch(`${API_URL}/ws/get-rooms`, {
        method: "GET",
      });

      const data = await res.json();
      if (res.ok) {
        setRooms(data);
      }
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    getRooms();
  }, []);

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault();

    try {
      const res = await fetch(`${API_URL}/ws/create-room`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          id: uuidv4(),
          name: roomName,
          password: "secret",
        }),
      });

      const data = await res.json();
      if (res.ok) {
        getRooms();
        setRoomName("");
      }
    } catch (err) {
      console.log(err);
    }
  };

  const joinRoom = (roomId: string) => {
    const ws = new WebSocket(
      `${WEBSOCKET_URL}/ws/join-room/${roomId}?userId=${user.id}&username=${user.username}&password=secret`
    );
    if (ws.OPEN) {
      setConn(ws);
      router.push("/app");
      return;
    }
  };

  return (
    <div className="my-8 px-0 md:mx-20 w-full h-full">
      <div className="flex justify-center mt-3 p-5">
        <input
          type="text"
          className="border border-grey p-2 rounded-md focus:outline-none focus:border-blue"
          placeholder="Room Name"
          value={roomName}
          onChange={(e) => setRoomName(e.target.value)}
        />
        <button
          className="bg-black border text-white rounded-md p-2 md:ml-4"
          onClick={submitHandler}
        >
          Create Room
        </button>
      </div>
      <div className="mt-6">
        <div className="font-bold">Available Rooms</div>
        <div className="grid grid-cols-1 md:grid-cols-5 gap-4 mt-6">
          {rooms.map((room, index) => (
            <div
              key={index}
              className="border border-black p-4 flex items-center rounded-md w-full"
            >
              <div className="w-full">
                <div className="text-sm">Room</div>
                <div className="text-black font-bold text-lg">{room.name}</div>
              </div>
              <div className="">
                <button
                  className="px-4 text-white bg-black rounded-md"
                  onClick={() => joinRoom(room.id)}
                >
                  Join
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default page;
