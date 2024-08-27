import Table from './Table'
import { render } from '@testing-library/react'

const mockData = [
  {
    id: 1,
    type: 'Unauthorized Access',
    premise: 'S3 Tower1',
    cameraId: 'CCTV1',
    date: '20/10/2023'
  },
  {
    id: 2,
    type: 'Unauthorized Access',
    premise: 'S3 Tower3',
    cameraId: 'CCTV12',
    date: '20/10/2023'
  },
  {
    id: 3,
    type: 'Unauthorized Access',
    premise: 'S3 Tower3',
    cameraId: 'CCTV13',
    date: '20/10/2023'
  }
]

const mockColumns = [
  {
    Header: 'ID',
    id: 'id',
    accessor: 'id'
  },
  {
    Header: 'Type',
    id: 'type',
    alignRight: true,
    sortable: false,
    accessor: 'type'
  }
]

const mockColumnsWithCell = [
  {
    Header: 'ID',
    id: 'id',
    accessor: 'id'
  },
  {
    Header: 'Type',
    id: 'type',
    alignRight: true,
    sortable: false,
    accessor: 'type',
    Cell: jest.fn(),
    Header: jest.fn()
  }
]

const pageCount = 10

describe('Table Component', () => {
  test('render Table component successfully with data', () => {
    const { getAllByText } = render(
      <Table columns={mockColumns} pageCount={pageCount} page={0} data={mockData} />
    )
    expect(getAllByText('Unauthorized Access')[1]).toBeTruthy()
  })

  test('render Table component successfully', () => {
    const { getByText } = render(
      <Table columns={mockColumns} pageCount={pageCount} page={0} />
    )
    expect(getByText('No records found')).toBeTruthy()
  })

  test('render Table component successfully with Cell and Header', () => {
    const { getAllByText } = render(
      <Table columns={mockColumnsWithCell} pageCount={pageCount} page={0} data={mockData} />
    )
    expect(getAllByText('Loading...')[0]).toBeTruthy()
  })

  test('render Table component with loading', () => {
    const { getAllByText } = render(
      <Table columns={mockColumnsWithCell} pageCount={pageCount} page={0} data={mockData} loading />
    )
    expect(getAllByText('Loading...')[0]).toBeTruthy()
  })
})
