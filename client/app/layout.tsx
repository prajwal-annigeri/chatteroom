import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import AuthContextProvider from "../modules/auth_provider";
import WebSocketProvider from "../modules/websocket_provider";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Chatteroom",
  description: "Join a room and start chatting",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html>
      <AuthContextProvider>
        <WebSocketProvider>
          <body className={inter.className}>{children}</body>
        </WebSocketProvider>
      </AuthContextProvider>
    </html>
  );
}
