import z from 'zod'

const { string, object } = z

const requestSchema = object({
  assignee: object({
    label: string(),
    value: string().min(1, 'Assignee is a required field')
  })
})

export default requestSchema
