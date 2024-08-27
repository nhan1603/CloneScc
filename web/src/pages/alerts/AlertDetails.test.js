import { render } from '@testing-library/react'
import { screen } from '@testing-library/dom'
import userEvent from '@testing-library/user-event'
import { BrowserRouter } from 'react-router-dom'
import AlertDetails from './AlertDetails'

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
      }
    },
    isLoading: false,
    error: false,
  })
})

describe('AlertDetails Component', () => {
  test('render AlertDetails with default value', () => {
    const { getByText, getAllByText } = render(
      <BrowserRouter>
        <AlertDetails />
      </BrowserRouter>
    )
    expect(getAllByText('Create Request')[0]).toBeTruthy()
    const createRequest = screen.getAllByTestId('create-request-details')[0]
    userEvent.click(createRequest)
    expect(getByText('Camera')).toBeTruthy()
  })
})
