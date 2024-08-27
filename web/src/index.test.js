import ReactDOM from 'react-dom/client'

jest.mock('react-dom', () => {
  return {
    ...jest.requireActual('react-dom'),
    createRoot: () => {
      return {
        render: jest.fn()
      }
    }
  }
})

describe('Application root', () => {
  it('should render without crashing', () => {
    const root = ReactDOM.createRoot(document.getElementById('root'))
    require('./index.js')
    expect(root.render).toBeDefined()
  })
})
