"use client";

import React, { useState, useRef, useContext, useEffect } from "react";
import { API_URL } from "../../constants";
import Chat from "../../components/chat";
import { WebSocketContext } from "../../modules/websocket_provider";
import { useRouter } from "next/navigation";
import autosize from "autosize";
import { AuthContext } from "../../modules/auth_provider";

export type Message = {
  content: string;
  client_id: string;
  username: string;
  room_id: string;
  type: "recv" | "self";
};
const page = () => {
  const [messages, setMessages] = useState<Array<Message>>([]);
  const textarea = useRef<HTMLTextAreaElement>(null);
  const { conn } = useContext(WebSocketContext);
  const [users, setUsers] = useState<Array<{ username: string }>>([]);
  const { user } = useContext(AuthContext);

  const router = useRouter();

  useEffect(() => {
    if (conn === null) {
      router.push("/");
      return;
    }

    const roomId = conn.url.split("/"[5]);
    async function getUsers() {
      try {
        const res = await fetch(`${API_URL}/ws/get-clients/${roomId}`, {
          method: "GET",
          headers: { "Content-Type": "application/json" },
        });

        const data = await res.json();
        setUsers(data);
      } catch (err) {
        console.log(err);
      }
    }
    getUsers();
  }, []);
  useEffect(() => {
    if (textarea.current) {
      autosize(textarea.current);
    }

    if (conn === null) {
      router.push("/");
      return;
    }

    conn.onmessage = (message) => {
      const msg: Message = JSON.parse(message.data);
      if (msg.content === "A new user has joined the room") {
        setUsers([...users, { username: msg.username }]);
      }

      if (msg.content.endsWith("has left the room")) {
        const deleteUser = users.filter(
          (user) => user.username != msg.username
        );
        setUsers([...deleteUser]);
        setMessages([...messages, msg]);
        return;
      }

      user.username === msg.username
        ? (msg.type = "self")
        : (msg.type = "recv");
      setMessages([...messages, msg]);
    };

    conn.onclose = () => {};
    conn.onerror = () => {};
    conn.onopen = () => {};
  }, [textarea, messages, conn, users]);

  const sendMessage = () => {
    if (!textarea.current?.value) return;
    if (conn === null) {
      router.push("/");
      return;
    }

    conn.send(textarea.current.value);
    textarea.current.value = "";
  };
  return (
    <>
      <div className="flex flex-col w-full">
        <div className="p-4 md:mx-6 mb-14">
          <Chat data={messages} />
        </div>
        <div className="fixed bottom-0 mt-4 w-full">
          <div className="flex md:flex-row px-4 py-2 bg-grey md:mx-4 rounded-md">
            <div className="flex w-full mr-4 rounded-md border border-blue">
              <textarea
                ref={textarea}
                placeholder="Enter message"
                className="w-full h-10 p-2 rounded-md foxus:outline-none"
                style={{ resize: "none" }}
              />
            </div>
            <div className="flex items-center">
              <button
                className="p-2 rounded-md bg-blue text-white"
                onClick={sendMessage}
              >
                Send
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default page;
