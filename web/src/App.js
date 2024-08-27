import './App.scss'
import {
  Outlet,
  Route,
  RouterProvider,
  createBrowserRouter,
  createRoutesFromElements
} from 'react-router-dom'
import Alerts from 'pages/alerts'
import Dashboard from 'pages/dashboard'
import ProtectedRoute from 'ProtectedRoute'
import Login from 'pages/login'
import Toastify from 'components/toastify'
import { RequestManagement, RequestManagementDetails } from 'pages/requests-management'
import AlertDetails from 'pages/alerts/AlertDetails'
import useWebSocket from 'react-use-websocket'
import { addToast } from 'components/toastify/Toast'
import { formatDateTime } from 'utils/datetime'
import { getToken } from 'data/auth'
import { setAlertMessages, getAlertMessages } from 'data/alertMessage'
import { getFloor } from 'utils'
import ErrorBoundary from './ErrorBoundary'
import usePremiseParam from 'hooks/usePremiseParam'
import { useEffect } from 'react'

const IndexEl = () => {
  const { lastMessage } = useWebSocket(process.env.REACT_APP_SOCKET_API_ENDPOINT)
  const currentMessage = lastMessage?.data && JSON.parse(lastMessage?.data)
  const alertMessages = getAlertMessages()
  const { withPremiseParam } = usePremiseParam()

  const toastNoti = () => {
    return addToast({
      title: currentMessage?.cctv_name + ' - ' + getFloor(currentMessage?.floor_number),
      message: currentMessage?.description,
      link: withPremiseParam(currentMessage?.id ? `/alerts/${currentMessage.id}` : '/alerts/'),
      status: currentMessage?.type,
      date: `${formatDateTime(currentMessage?.incident_at)}`,
      options: {
        toastId: currentMessage?.id ?? new Date().getTime()
      }
    })
  }

  useEffect(() => {
    if (currentMessage) {
      let allMessages = []
      if (alertMessages?.length > 0 && currentMessage?.id !== alertMessages?.[0].id) {
        allMessages = [currentMessage, ...alertMessages]
      } else {
        allMessages.push(currentMessage)
      }
      setAlertMessages(allMessages)
      getToken() && toastNoti()
    }
  }, [JSON.stringify(currentMessage)])

  return <Outlet />
}

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path="/" errorElement={<ErrorBoundary />} element={<IndexEl />}>
      <Route
        index
        element={
          <ProtectedRoute>
            <Dashboard />
          </ProtectedRoute>
        }
      />
      <Route path="alerts">
        <Route
          index
          element={
            <ProtectedRoute>
              <Alerts />
            </ProtectedRoute>
          }
        />
        <Route
          path=":alertId"
          element={
            <ProtectedRoute>
              <AlertDetails />
            </ProtectedRoute>
          }
        />
      </Route>
      <Route path="requests">
        <Route
          index
          element={
            <ProtectedRoute>
              <RequestManagement />
            </ProtectedRoute>
          }
        />
        <Route
          index
          path=":requestId"
          element={
            <ProtectedRoute>
              <RequestManagementDetails />
            </ProtectedRoute>
          }
        />
      </Route>
      <Route path="login" element={<Login />} />
      <Route path="*" element={<ErrorBoundary error={new Error('Page not found')} />} />
    </Route>
  )
)

function App() {
  return (
    <>
      <Toastify />
      <RouterProvider router={router} />
    </>
  )
}

export default App
