export const parseJwt = (token) => {
  try {
    return JSON.parse(atob(token.split('.')[1]))
  } catch (e) {
    return null
  }
}

const isNullOrUndefined = (value) => value === null || value === undefined

export const getFloor = (floorNumber) => {
  if (isNullOrUndefined(floorNumber)) {
    return floorNumber
  }
  return Number(floorNumber) >= 1 ? `Floor ${floorNumber}` : 'G-Floor'
}

const videoExtensions = ['.mp4', '.m4p', '.mpg', '.mp2', '.mpeg', '.mpe', '.mpv', '.m4v']

export const checkIsVideoOfUrl = (extension) => {
  return videoExtensions.find((element) => element === extension)
}
