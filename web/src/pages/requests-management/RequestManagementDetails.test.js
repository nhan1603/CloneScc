import { render } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import userEvent from '@testing-library/user-event'
import RequestManagementDetails from './RequestManagementDetails'

jest.mock('hooks/useGet', () => {
  return {
    __esModule: true,
    default: () => {
      return {
        response: {
          data: {
            id: '31116122895990789',
            title: 'Camera 1-31115933061791749',
            author: 'Hue',
            assignee: 'John Henry',
            message: '',
            startTime: '2023-08-03T15:50:54.528213Z',
            alertDetail: {
              id: '31115933061791749',
              type: 'Suspicious Activities',
              premiseName: 'Sunrise Tower',
              premiseLocation: '307/12 Nguyen Van Troi St, W1, Tan Binh',
              cctvDevice: 'Camera 1',
              cctvDeviceFloor: 1,
              isAcknowledged: false,
              incidentAt: '2023-08-03T15:49:01.374202Z'
            },
            respond: null
          }
        }
      }
    }
  }
})

describe('RequestManagementDetails Component', () => {
  test('render RequestManagementDetails with default value', () => {
    const { getByText, getAllByText, queryByText } = render(
      <BrowserRouter>
        <RequestManagementDetails />
      </BrowserRouter>
    )
    expect(getByText('Security Command Center')).toBeTruthy()
    expect(getByText('Request Management')).toBeTruthy()

    // Click to show Create Request modal
    userEvent.click(getAllByText('Create Request')[0])
    expect(getByText('Detailed Information')).toBeTruthy()

    // Click to close Create Request modal
    userEvent.click(getAllByText('Cancel')[0])
    const fieldModal = queryByText('Detailed Information')
    expect(fieldModal).not.toBeInTheDocument()

    expect(getByText('Hue')).toBeInTheDocument()
    expect(getByText('John Henry')).toBeInTheDocument()
    expect(getByText('Suspicious Activities')).toBeInTheDocument()
    expect(getByText('Sunrise Tower')).toBeInTheDocument()
    expect(getByText('Camera 1 - Floor 1')).toBeInTheDocument()
  })
})
