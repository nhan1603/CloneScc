export const setAlertMessages = (alertMessages) => {
  alertMessages && localStorage.setItem('alertMessages', JSON.stringify(alertMessages))
}

export const getAlertMessages = () => {
  const messages = localStorage.getItem('alertMessages')
  return messages && JSON.parse(messages)
}

export const removeAlertMessages = () => {
  localStorage.removeItem('alertMessages')
}
