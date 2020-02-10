import Redis from 'redis'

const client = Redis.createClient()
client.on('ready', () => console.info('Ready'))
client.on('error', e => console.error(`Error: ${e}`))
client.on('reconnecting', params =>  console.info('Reconnect'))

// type RedisClient = {
//   on(event: 'ready', f: () => void): void
//   on(event: 'error', f: (e: Error) => void): void
//   on(event: 'reconnecting', f: (params: { attempt: number, delay: number }) => void): void
// }

type Events = {
  ready: void
  error: Error
  reconnecting: { attempt: number, delay: number }
}

type RedisClient = {
  on<E extends keyof Events>(event: E, f: (arg: Events[E]) => void): void
  emit<E extends keyof Events>(event: E, f: (arg: Events[E]) => void): void
}
