// const puppeteer = require('puppeteer');
import Puppeteer from 'puppeteer'
import DayJS from 'dayjs'

(async () => {
  const browser = await Puppeteer.launch(
    { headless: true }
  )
  const url = process.env.URL || ''
  const page = await browser.newPage()
  await page.goto(url)
  const items = await page.$$('tr.default')
  const data = await Promise.all(
    items.map(
      async (item) => {
        const itemColumns = await item.$$('td')
        const categoryElement = await itemColumns[0].$('a')
        const categoryLink = categoryElement && await (await categoryElement.getProperty('href')).jsonValue()
        const category = typeof categoryLink === 'string' && new URL(categoryLink).search.replace(/\?c\=/, '')
        const nameElement = await itemColumns[1].$('td > a')
        const nameLink = nameElement && await (await nameElement.getProperty('href')).jsonValue()
        const nameText = nameElement && await (await nameElement.getProperty('title')).jsonValue()

        const sizeElement = itemColumns[3]

        const size = await (await sizeElement.getProperty('textContent')).jsonValue()

        const dateElement = itemColumns[4]
        const createdAt = await (await dateElement.getProperty('textContent')).jsonValue()
        const created = typeof createdAt === 'string' && DayJS(createdAt).format()
        const seedersElement = itemColumns[5]
        const seeders = Number(await (await seedersElement.getProperty('textContent')).jsonValue())
        const leechersElement = itemColumns[6]
        const leechers = Number(await (await leechersElement.getProperty('textContent')).jsonValue())
        const completedDownloadsElement = itemColumns[7]
        const completedDownloads = Number(await (await completedDownloadsElement.getProperty('textContent')).jsonValue())
        return {
          category,
          nameLink,
          nameText,
          size,
          created,
          seeders,
          leechers,
          completedDownloads
        }
      }
    )
  )

  console.log(data)

  await page.close()
  await browser.close()
})();
