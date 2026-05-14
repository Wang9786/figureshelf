import { Navigate, Route, Routes } from 'react-router-dom'
import LoginPage from './pages/LoginPage'
import RegisterPage from './pages/RegisterPage'
import DashboardPage from './pages/DashboardPage'
import FigureListPage from './pages/FigureListPage'
import FigureCreatePage from './pages/FigureCreatePage'
import FigureDetailPage from './pages/FigureDetailPage'
import FigureEditPage from './pages/FigureEditPage'
import UpcomingPaymentsPage from './pages/UpcomingPaymentsPage'
import UpcomingReleasesPage from './pages/UpcomingReleasesPage'

function App() {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/login" replace />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/dashboard" element={<DashboardPage />} />
      <Route path="/figures" element={<FigureListPage />} />
      <Route path="/figures/new" element={<FigureCreatePage />} />
      <Route path="/figures/upcoming-payments" element={<UpcomingPaymentsPage />} />
      <Route path="/figures/upcoming-releases" element={<UpcomingReleasesPage />} />
      <Route path="/figures/:id" element={<FigureDetailPage />} />
      <Route path="/figures/:id/edit" element={<FigureEditPage />} />
    </Routes>
  )
}

export default App