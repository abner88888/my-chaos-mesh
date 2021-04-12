import * as d3 from 'd3'

import day, { format } from 'lib/dayjs'

import { Event } from 'api/events.type'
import { IntlShape } from 'react-intl'
import { Theme } from 'slices/settings'
import _debounce from 'lodash.debounce'
import { truncate } from '../utils'
import wrapText from './wrapText'

export default function gen({
  root,
  events,
  intl,
  theme,
  options = {
    enableLegends: true,
  },
}: {
  root: HTMLElement
  events: Event[]
  intl: IntlShape
  theme: Theme
  options?: {
    enableLegends?: boolean
    onSelectEvent?: (e: Event) => () => void
  }
}) {
  const { enableLegends, onSelectEvent } = options

  let width = root.offsetWidth
  const height = root.offsetHeight

  const margin = {
    top: 0,
    right: 0,
    bottom: 30,
    left: 0,
  }
  updateMargin()

  function updateMargin() {
    margin.right = enableLegends && document.documentElement.offsetWidth > 768 ? 150 : 0
  }

  const svg = d3
    .select(root)
    .append('svg')
    .attr('class', theme === 'light' ? 'chaos-chart' : 'chaos-chart-dark')
    .attr('width', width)
    .attr('height', height)

  const halfHourLater = (events.length ? day(events[events.length - 1].start_time) : day()).add(0.5, 'h')

  const x = d3
    .scaleLinear()
    .domain([halfHourLater.subtract(1, 'h'), halfHourLater])
    .range([margin.left, width - margin.right])
  const xAxis = d3
    .axisBottom(x)
    .ticks(6)
    .tickFormat(d3.timeFormat('%m-%d %H:%M') as (dv: Date | { valueOf(): number }, i: number) => string)
  const gXAxis = svg
    .append('g')
    .attr('class', 'axis')
    .attr('transform', `translate(0, ${height - margin.bottom})`)
    .call(xAxis)

  // Wrap long text, also used in zoom() and reGen()
  svg.selectAll('.tick text').call(wrapText, 30)

  const colorPalette = d3
    .scaleOrdinal<string, string>()
    .domain(events.map((d) => d.experiment_id))
    .range(d3.schemeTableau10)

  const allUniqueExperiments = [...new Set(events.map((d) => d.experiment + '/' + d.experiment_id))].map((d) => {
    const [name, uuid] = d.split('/')

    return {
      name,
      uuid,
    }
  })
  const allUniqueUUIDs = allUniqueExperiments.map((d) => d.uuid)

  const y = d3
    .scaleBand()
    .domain(allUniqueUUIDs)
    .range([0, height - margin.top - margin.bottom])
    .padding(0.5)

  const timelines = svg
    .append('g')
    .attr('transform', `translate(${margin.left}, ${margin.top})`)
    .attr('stroke-opacity', 0.12)
    .selectAll()
    .data(allUniqueUUIDs)
    .join('line')
    .attr('y1', (d) => y(d)! + y.bandwidth() / 2)
    .attr('y2', (d) => y(d)! + y.bandwidth() / 2)
    .attr('x2', width - margin.right - margin.left)
    .attr('stroke', colorPalette)

  // clipX
  svg
    .append('clipPath')
    .attr('id', 'clip-x-axis')
    .append('rect')
    .attr('x', margin.left)
    .attr('y', 0)
    .attr('width', width - margin.left - margin.right)
    .attr('height', height - margin.bottom)
  const gMain = svg.append('g').attr('clip-path', 'url(#clip-x-axis)')

  // legends
  const legendsRoot = d3.select(document.createElement('div')).attr('class', 'chaos-events-legends')
  if (enableLegends) {
    legends()
  }
  function legends() {
    const legends = legendsRoot
      .selectAll()
      .data(allUniqueExperiments)
      .enter()
      .append('div')
      .on('click', function (d) {
        const _events = events.filter((e) => e.experiment_id === d.uuid)
        const event = _events[_events.length - 1]

        svg
          .transition()
          .duration(750)
          .call(
            zoom.transform,
            d3.zoomIdentity
              .translate((width - margin.left - margin.right) / 2, 0)
              .scale(3)
              .translate(-x(day(event.start_time))!, 0)
          )
      })
    legends
      .append('div')
      .attr('class', 'square')
      .attr('style', (d) => `background: ${colorPalette(d.uuid)};`)
    legends
      .insert('div')
      .attr('class', 'experiment')
      .attr('title', (d) => d.name)
      .text((d) => truncate(d.name))
  }

  // event circles
  const circles = gMain
    .selectAll()
    .data(events)
    .enter()
    .append('circle')
    .attr('cx', (d) => x(day(d.start_time))!)
    .attr('cy', (d) => y(d.experiment_id)! + y.bandwidth() / 2 + margin.top)
    .attr('r', 4)
    .attr('fill', (d) => colorPalette(d.experiment_id))
    .style('cursor', 'pointer')

  const zoom = d3.zoom<SVGSVGElement, unknown>().scaleExtent([0.1, 6]).on('zoom', zoomed)
  function zoomed() {
    const eventTransform = d3.event.transform

    const newX = eventTransform.rescaleX(x)

    gXAxis.call(xAxis.scale(newX))
    svg.selectAll('.tick text').call(wrapText, 30)
    circles.attr('cx', (d) => newX(day(d.start_time)))
  }
  svg.call(zoom as any)

  const tooltip = d3
    .select(document.createElement('div'))
    .attr('class', 'chaos-event-tooltip')
    .call(createTooltip as any)

  function createTooltip(el: d3.Selection<HTMLElement, any, any, any>) {
    el.style('position', 'absolute')
      .style('top', 0)
      .style('left', 0)
      .style('padding', '0.25rem 0.75rem')
      .style('background', theme === 'light' ? '#fff' : 'rgba(0, 0, 0, 0.54)')
      .style('font', '1rem')
      .style('border', `1px solid ${theme === 'light' ? 'rgba(0, 0, 0, 0.12)' : 'rgba(255, 255, 255, 0.12)'}`)
      .style('border-radius', '4px')
      .style('opacity', 0)
      .style('transition', 'top 0.25s ease, left 0.25s ease')
      .style('z-index', 999)
  }

  function genTooltipContent(d: Event) {
    return `<b>${intl.formatMessage({ id: 'events.event.experiment' })}: ${d.experiment}</b>
            <br />
            <b>
              ${intl.formatMessage({ id: 'common.status' })}: ${
      d.finish_time
        ? intl.formatMessage({ id: 'experiments.state.finished' })
        : intl.formatMessage({ id: 'experiments.state.running' })
    }
            </b>
            <br />
            <br />
            <span style="color: ${theme === 'light' ? 'rgba(0, 0, 0, 0.54)' : '#fff'};">
              ${intl.formatMessage({ id: 'events.event.started' })}: ${format(d.start_time)}
            </span>
            <br />
            ${
              d.finish_time
                ? `<span style="color: ${theme === 'light' ? 'rgba(0, 0, 0, 0.54)' : '#fff'};">${intl.formatMessage({
                    id: 'events.event.ended',
                  })}: ${format(d.finish_time)}</span>`
                : ''
            }
            `
  }

  circles
    .on('click', function (d) {
      if (typeof onSelectEvent === 'function') {
        onSelectEvent(d)()
      }

      svg
        .transition()
        .duration(750)
        .call(
          zoom.transform as any,
          d3.zoomIdentity
            .translate(width / 2, 0)
            .scale(2)
            .translate(-x(day(d.start_time))!, 0)
        )
    })
    .on('mouseover', function (d) {
      let [x, y] = d3.mouse(this)

      tooltip.html(genTooltipContent(d))
      const { width } = tooltip.node()!.getBoundingClientRect()

      y += 50
      if (x > (root.offsetWidth / 3) * 2) {
        x -= width
      }
      if (y > (root.offsetHeight / 3) * 2) {
        y -= 200
      }

      tooltip
        .style('left', x + 'px')
        .style('top', y + 'px')
        .style('opacity', 1)
    })
    .on('mouseleave', function () {
      tooltip.style('opacity', 0)
    })

  function reGen() {
    const newWidth = root.offsetWidth
    width = newWidth

    updateMargin()

    svg.attr('width', width).call(zoom.transform, d3.zoomIdentity)
    gXAxis.call(xAxis.scale(x.range([margin.left, width - margin.right])))
    svg.selectAll('.tick text').call(wrapText, 30)
    timelines.attr('x2', width - margin.right - margin.left)
    circles.attr('x', (d) => x(day(d.start_time))!)
  }

  d3.select(window).on('resize', _debounce(reGen, 250))

  root.appendChild(legendsRoot.node()!)
  root.appendChild(tooltip.node()!)
  root.style.position = 'relative'
}
