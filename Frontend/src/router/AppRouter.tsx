import { Route,Routes } from "react-router-dom"
import AuthRoutes from "../auth/routes/AuthRoutes"
import HomeRoutes from "../home/routes/HomeRoutes"

export const AppRouter = () => {
  return (
    <Routes>
        {/* Login */}
        <Route path="auth/*" element={<AuthRoutes />}/>
    
        {/* Home */}
        <Route path="/*" element={<HomeRoutes />}/>
    </Routes>
  )
}
