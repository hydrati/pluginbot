export default (ms: number, ...args: any[]) => (
  new Promise(r => setTimeout(r, ms, ...args))
)