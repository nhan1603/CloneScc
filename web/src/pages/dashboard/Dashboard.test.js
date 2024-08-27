import { render } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import Dashboard from './Dashboard'

jest.mock('hooks/useGet', () => {
  return {
    __esModule: true,
    default: (url) => {
      if (url === '/api/authenticated/v1/devices') {
        return {
          response: {
            data: {
              items: [
                {
                  id: 30,
                  premiseID: 7,
                  premiseName: 'Lua Office',
                  premiseLocation: '40 Nguyen Van Troi st., W1, Tan Binh, HCMC',
                  deviceName: 'Camera 1',
                  deviceCode: 'cctv_cam30',
                  isActive: true,
                  floorNumber: 0,
                  deviceURL: ''
                },
                {
                  id: 20,
                  premiseID: 5,
                  premiseName: 'Sunrise Tower',
                  premiseLocation: '307/12 Nguyen Van Troi St, W1, Tan Binh',
                  deviceName: 'Camera 1',
                  deviceCode: 'cctv_cam1',
                  isActive: true,
                  floorNumber: 1,
                  deviceURL: ''
                },
                {
                  id: 31,
                  premiseID: 7,
                  premiseName: 'Lua Office',
                  premiseLocation: '40 Nguyen Van Troi st., W1, Tan Binh, HCMC',
                  deviceName: 'Camera 2',
                  deviceCode: 'cctv_cam31',
                  isActive: true,
                  floorNumber: 1,
                  deviceURL: ''
                },
                {
                  id: 21,
                  premiseID: 5,
                  premiseName: 'Sunrise Tower',
                  premiseLocation: '307/12 Nguyen Van Troi St, W1, Tan Binh',
                  deviceName: 'Camera 2',
                  deviceCode: 'cctv_cam2',
                  isActive: true,
                  floorNumber: 1,
                  deviceURL: ''
                },
                {
                  id: 22,
                  premiseID: 5,
                  premiseName: 'Sunrise Tower',
                  premiseLocation: '307/12 Nguyen Van Troi St, W1, Tan Binh',
                  deviceName: 'Camera 3',
                  deviceCode: 'cctv_cam3',
                  isActive: true,
                  floorNumber: 2,
                  deviceURL: ''
                },
                {
                  id: 32,
                  premiseID: 7,
                  premiseName: 'Lua Office',
                  premiseLocation: '40 Nguyen Van Troi st., W1, Tan Binh, HCMC',
                  deviceName: 'Camera 3',
                  deviceCode: 'cctv_cam32',
                  isActive: true,
                  floorNumber: 1,
                  deviceURL: ''
                },
                {
                  id: 33,
                  premiseID: 7,
                  premiseName: 'Lua Office',
                  premiseLocation: '40 Nguyen Van Troi st., W1, Tan Binh, HCMC',
                  deviceName: 'Camera 4',
                  deviceCode: 'cctv_cam33',
                  isActive: true,
                  floorNumber: 2,
                  deviceURL: ''
                },
                {
                  id: 23,
                  premiseID: 5,
                  premiseName: 'Sunrise Tower',
                  premiseLocation: '307/12 Nguyen Van Troi St, W1, Tan Binh',
                  deviceName: 'Camera 4',
                  deviceCode: 'cctv_cam4',
                  isActive: true,
                  floorNumber: 3,
                  deviceURL: ''
                },
                {
                  id: 24,
                  premiseID: 6,
                  premiseName: 'Bitexco Financial Tower',
                  premiseLocation: '2 Hai Ba Trung St, Ben Nghe, District 1',
                  deviceName: 'Camera 5',
                  deviceCode: 'cctv_cam5',
                  isActive: true,
                  floorNumber: 1,
                  deviceURL: ''
                },
                {
                  id: 34,
                  premiseID: 7,
                  premiseName: 'Lua Office',
                  premiseLocation: '40 Nguyen Van Troi st., W1, Tan Binh, HCMC',
                  deviceName: 'Camera 5',
                  deviceCode: 'cctv_cam34',
                  isActive: true,
                  floorNumber: 2,
                  deviceURL: ''
                },
                {
                  id: 35,
                  premiseID: 7,
                  premiseName: 'Lua Office',
                  premiseLocation: '40 Nguyen Van Troi st., W1, Tan Binh, HCMC',
                  deviceName: 'Camera 6',
                  deviceCode: 'cctv_cam35',
                  isActive: true,
                  floorNumber: 3,
                  deviceURL: ''
                },
                {
                  id: 25,
                  premiseID: 6,
                  premiseName: 'Bitexco Financial Tower',
                  premiseLocation: '2 Hai Ba Trung St, Ben Nghe, District 1',
                  deviceName: 'Camera 6',
                  deviceCode: 'cctv_cam6',
                  isActive: true,
                  floorNumber: 2,
                  deviceURL: ''
                },
                {
                  id: 36,
                  premiseID: 7,
                  premiseName: 'Lua Office',
                  premiseLocation: '40 Nguyen Van Troi st., W1, Tan Binh, HCMC',
                  deviceName: 'Camera 7',
                  deviceCode: 'cctv_cam36',
                  isActive: true,
                  floorNumber: 4,
                  deviceURL: ''
                },
                {
                  id: 40,
                  premiseID: 8,
                  premiseName: 'Head Office',
                  premiseLocation: '10 Truong Dinh st., W8, District 3, HCMC',
                  deviceName: 'CCTV 1',
                  deviceCode: 'cctv_cam40',
                  isActive: true,
                  floorNumber: 0,
                  deviceURL: ''
                },
                {
                  id: 49,
                  premiseID: 8,
                  premiseName: 'Head Office',
                  premiseLocation: '10 Truong Dinh st., W8, District 3, HCMC',
                  deviceName: 'CCTV 10',
                  deviceCode: 'cctv_cam49',
                  isActive: true,
                  floorNumber: 7,
                  deviceURL: ''
                },
                {
                  id: 41,
                  premiseID: 8,
                  premiseName: 'Head Office',
                  premiseLocation: '10 Truong Dinh st., W8, District 3, HCMC',
                  deviceName: 'CCTV 2',
                  deviceCode: 'cctv_cam41',
                  isActive: true,
                  floorNumber: 1,
                  deviceURL: ''
                },
                {
                  id: 42,
                  premiseID: 8,
                  premiseName: 'Head Office',
                  premiseLocation: '10 Truong Dinh st., W8, District 3, HCMC',
                  deviceName: 'CCTV 3',
                  deviceCode: 'cctv_cam42',
                  isActive: true,
                  floorNumber: 1,
                  deviceURL: ''
                },
                {
                  id: 43,
                  premiseID: 8,
                  premiseName: 'Head Office',
                  premiseLocation: '10 Truong Dinh st., W8, District 3, HCMC',
                  deviceName: 'CCTV 4',
                  deviceCode: 'cctv_cam43',
                  isActive: true,
                  floorNumber: 2,
                  deviceURL: ''
                },
                {
                  id: 44,
                  premiseID: 8,
                  premiseName: 'Head Office',
                  premiseLocation: '10 Truong Dinh st., W8, District 3, HCMC',
                  deviceName: 'CCTV 5',
                  deviceCode: 'cctv_cam44',
                  isActive: true,
                  floorNumber: 2,
                  deviceURL: ''
                },
                {
                  id: 45,
                  premiseID: 8,
                  premiseName: 'Head Office',
                  premiseLocation: '10 Truong Dinh st., W8, District 3, HCMC',
                  deviceName: 'CCTV 6',
                  deviceCode: 'cctv_cam45',
                  isActive: true,
                  floorNumber: 3,
                  deviceURL: ''
                }
              ],
              pagination: {
                totalCount: 23,
                currentPage: 1,
                limit: 20
              }
            }
          }
        }
      } else if (url === '/api/authenticated/v1/premises') {
        return {
          response: {
            data: {
              items: [
                {
                  id: 6,
                  name: 'Bitexco Financial Tower',
                  location: '2 Hai Ba Trung St, Ben Nghe, District 1',
                  premises_code: 'P002',
                  description: 'Bitexco Financial Tower',
                  cctv_count: 2
                },
                {
                  id: 8,
                  name: 'Head Office',
                  location: '10 Truong Dinh st., W8, District 3, HCMC',
                  premises_code: 'P004',
                  description: 'Head Office',
                  cctv_count: 10
                },
                {
                  id: 7,
                  name: 'Lua Office',
                  location: '40 Nguyen Van Troi st., W1, Tan Binh, HCMC',
                  premises_code: 'P003',
                  description: 'Lua Office',
                  cctv_count: 7
                },
                {
                  id: 5,
                  name: 'Sunrise Tower',
                  location: '307/12 Nguyen Van Troi St, W1, Tan Binh',
                  premises_code: 'P001',
                  description: 'Sunrise Tower',
                  cctv_count: 4
                }
              ]
            }
          }
        }
      }
    }
  }
})

describe('Dashboard Component', () => {
  test('render Dashboard with default value', () => {
    const { getByText } = render(
      <BrowserRouter>
        <Dashboard />
      </BrowserRouter>
    )
    expect(getByText('Security Command Center')).toBeTruthy()
    expect(getByText('All')).toBeInTheDocument()
    expect(getByText('All premises')).toBeInTheDocument()
    expect(getByText('Lua Office - Camera 1')).toBeInTheDocument()
    expect(getByText('Sunrise Tower - Camera 1')).toBeInTheDocument()
    expect(getByText('Lua Office - Camera 2')).toBeInTheDocument()
    expect(getByText('Sunrise Tower - Camera 2')).toBeInTheDocument()
  })
})
