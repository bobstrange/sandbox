import createError from 'http-errors'
import { getEndedAuctions } from "../lib/getEndedAuctions"
import { closeAuction } from "../lib/cloceAuction"

async function processAuctions(_event, _context): Promise<{ closed: number }> {
  try {
    const auctionsToClose = await getEndedAuctions()
    const closeOperations = auctionsToClose.map(closeAuction)
    await Promise.all(closeOperations)
    return { closed: closeOperations.length }
  } catch (error) {
    console.error(error)
    throw new createError.InternalServerError(error)
  }
}

export const handler = processAuctions
