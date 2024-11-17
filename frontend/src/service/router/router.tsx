import MainLayout from "../../pages/layout";
import DashboardPage from "../../pages/dashboard";
import { createBrowserRouter, Navigate } from "react-router-dom";
import AuthLayout from "@/pages/auth/layout";
import LoginPage from "@/pages/auth/login";
import RegisterPage from "@/pages/auth/register";
import ForgotPasswordPage from "@/pages/auth/forgot-password";

const isAuthenticated = false;

// Protected Route Wrapper
const ProtectedRoute = ({ element }: {element: React.ReactNode}) => {
  return isAuthenticated ? element : <Navigate to="/login" replace />;
};

const authRoute = {
  path: "/",
  element: <AuthLayout />,
  children: [
    { path: "/login", element:<LoginPage /> },
    { path: "/register", element:<RegisterPage /> },
    { path: "/forgot-password", element:<ForgotPasswordPage /> },
  ]
}

export const router = createBrowserRouter([
  {
    path: "/",
    element: <MainLayout />,
    children:[
      { path: "/", element: <ProtectedRoute element={<DashboardPage />} /> },
      authRoute,
    ]
  },
]);