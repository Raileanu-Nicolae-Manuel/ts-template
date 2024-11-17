import { useState } from "react";
import { Outlet } from "react-router-dom";

export default function MainLayout () {
  const [dark, setDark] = useState(true);
  return (
    <main className="h-screen bg-slate-200 dark:bg-slate-800 text-slate-800 dark:text-slate-400">
      {/* TODO: add navbar */}
      <div className="h-full">
        <Outlet />
      </div>
    </main>
  );
}