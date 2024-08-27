import moment from 'moment'
export const formatDateTime = (value) => {
  return value ? moment(value).format('D MMM YYYY, h:mm A') : '-'
}
