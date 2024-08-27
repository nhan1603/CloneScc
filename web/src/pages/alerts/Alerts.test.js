import { render } from '@testing-library/react'
import { screen } from '@testing-library/dom'
import userEvent from '@testing-library/user-event'
import { BrowserRouter } from 'react-router-dom'
import Alerts from './Alerts'

jest.mock('hooks/usePremiseParam', () => {
  return () => ({
    premiseId: '123',
    withPremiseParam: () => {
      return () => {
        return '/hello'
      }
    }
  })
})

jest.mock('hooks/useGet', () => {
  return () => ({
    response: {
      data: {
        items : [
          {
            "id": "31188536396922885",
            "type": "Unauthorized Access 3315",
            "premiseName": "Head Office",
            "premiseLocation": "10 Truong Dinh st., W8, District 3, HCMC",
            "cctvDevice": "CCTV 7",
            "cctvDeviceFloor": 4,
            "isAcknowledged": false,
            "incidentAt": "2023-08-04T03:50:16.332786Z",
            "date": "2023-08-04T03:50:16.332786Z",
            "premise": {
                "name": "Head Office",
                "location": "10 Truong Dinh st., W8, District 3, HCMC"
            },
            "cameraId": "CCTV 7 -  Floor 4"
          },
          {
            "id": "31188226102312965",
            "type": "Suspicious Activities",
            "premiseName": "Sunrise Tower",
            "premiseLocation": "307/12 Nguyen Van Troi St, W1, Tan Binh",
            "cctvDevice": "Camera 1",
            "cctvDeviceFloor": 1,
            "isAcknowledged": false,
            "incidentAt": "2023-08-04T03:47:11.378674Z",
            "date": "2023-08-04T03:47:11.378674Z",
            "premise": {
                "name": "Sunrise Tower",
                "location": "307/12 Nguyen Van Troi St, W1, Tan Binh"
            },
            "cameraId": "Camera 1 -  Floor 1"
          }
        ]
      }
    },
    isLoading: false,
    error: false,
  })
})

describe('Alerts Component', () => {
  test('render Alerts with default value', () => {
    const { getByText, getAllByText } = render(
      <BrowserRouter>
        <Alerts />
      </BrowserRouter>
    )
    expect(getAllByText('Create Request')[0]).toBeTruthy()
    const createRequest = screen.getAllByTestId('create-request')[0]
    userEvent.click(createRequest)
    expect(getByText('Security Command Center')).toBeTruthy()
  })
})
