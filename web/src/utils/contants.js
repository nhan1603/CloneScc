export const MENU_ITEMS = [
  {
    name: 'Dashboard',
    path: '/'
  },
  {
    name: 'Alerts',
    path: '/alerts',
    isMultiPath: true
  },
  {
    name: 'Request Management',
    path: '/requests',
    isMultiPath: true
  }
]

export const LIMIT = 10
export const GET_PREMISES_URL = '/api/authenticated/v1/premises'
export const GET_DEVICES_URL = '/api/authenticated/v1/devices'
export const GET_ALERTS_URL = '/api/authenticated/v1/alerts'
export const REQUESTS_URL = '/api/authenticated/v1/requests'
export const USERS_URL = '/api/authenticated/v1/users'
