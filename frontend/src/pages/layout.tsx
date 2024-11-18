import { useWebSocket } from "@/hooks/use-web-socket";
import { useState } from "react";
import { Outlet } from "react-router-dom";

export default function MainLayout () {
  const [dark, setDark] = useState(true);
  const ws = useWebSocket();
  const handleSendMessage = () => {
    ws.sendMessage({ 
      type: 'test',
      data: 'Hello Server!'
    });
  };
  return (
    <main className="h-screen bg-slate-200 dark:bg-slate-800 text-slate-800 dark:text-slate-400">
      {/* TODO: add navbar */}
      <button onClick={handleSendMessage}>Send Message</button>
      <div className="h-full">
        <Outlet />
      </div>
    </main>
  );
}