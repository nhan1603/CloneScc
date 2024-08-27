import { render } from '@testing-library/react'
import { screen } from '@testing-library/dom'
import userEvent from '@testing-library/user-event'
import { BrowserRouter } from 'react-router-dom'
import Notification from './Notification'

jest.mock('react-use-websocket', () => {
  return () => ({
    lastMessage: {
      data: '{"cctv_name":"CCTV name","floor_number":"G Floor","description":"Description","type":"Alert"}'
    }
  })
})

jest.mock('data/alertMessage', () => {
  return {
    getAlertMessages: () => {
      return [
        {
          "cctv_name": "Cam8",
          "floor_number": "Floor 5",
          "type": "Property Damage",
          "description": "",
          "incident_at": "2023-08-03T14:02:00Z"
        },
        {
          "cctv_name": "Cam8",
          "floor_number": "Floor 5",
          "type": "Property Damage",
          "description": "",
          "incident_at": "2023-08-03T14:02:00Z"
        },
        {
          "cctv_name": "Cam8123",
          "floor_number": "Floor 5",
          "type": "Property Damage",
          "description": "",
          "incident_at": "2023-08-03T14:02:00Z"
        }
      ]
    },
    removeAlertMessages: () => jest.fn()
  }
})

jest.mock('data/auth', () => ({
  getToken: () => true
}))

test('render Notification', () => {
  const { getByText, container, queryByText } = render(
    <BrowserRouter>
      <Notification />
    </BrowserRouter>
  )
  const notification = container.getElementsByClassName('notification')
  expect(notification.length).toBe(1)
  const elementNotification = screen.getByTestId('notification')
  userEvent.click(elementNotification)
  expect(getByText('Cam8123 - G-Floor')).toBeTruthy()

  // Mark all as read
  const elementMarkAll = screen.getByTestId('mark-all')
  userEvent.click(elementMarkAll)
  const cameraName = queryByText('Cam8123 - G-Floor')
  expect(cameraName).not.toBeInTheDocument()
})
