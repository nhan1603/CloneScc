import { render } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import RequestManagement from './RequestManagement'

jest.mock('hooks/useGet', () => {
  return () => ({
    response: {
      data: {
        items : [
          {
            "id": "31190985753026566",
            "alertID": "31190767045238790",
            "alert": "CCTV 10-31190767045238790",
            "alertType": "Suspicious Activities",
            "premiseName": "Head Office",
            "premiseLocation": "10 Truong Dinh st., W8, District 3, HCMC",
            "author": "Hue",
            "assignee": "Security",
            "status": "New",
            "startTime": "2023-08-04T04:14:36.260814Z",
            "verifiedAt": null
          },
          {
            "id": "31190921479512070",
            "alertID": "31190855964483590",
            "alert": "Camera 5-31190855964483590",
            "alertType": "Suspicious Activities",
            "premiseName": "Bitexco Financial Tower",
            "premiseLocation": "2 Hai Ba Trung St, Ben Nghe, District 1",
            "author": "Hai",
            "assignee": "My SG",
            "status": "New",
            "startTime": "2023-08-04T04:13:57.951158Z",
            "verifiedAt": null
          }
        ]
      }
    },
    isLoading: false,
    error: false,
  })
})

describe('RequestManagement Component', () => {
  test('render RequestManagement with default value', () => {
    const { getByText } = render(
      <BrowserRouter>
        <RequestManagement />
      </BrowserRouter>
    )
    expect(getByText('All premises')).toBeTruthy()
  })
})
