import { Line, mixins } from 'vue-chartjs'
const { reactiveProp } = mixins

export default {
  extends: Line,
  mixins: [reactiveProp],
  props: ['options'],
  mounted () {
    // this.chartData is created in the mixin.
    // If you want to pass options please create a local options object
    this.renderChart(this.chartData, this.options)
  }
}

/*

import { Line, mixins } from 'vue-chartjs'
const { reactiveProp } = mixins

export default {
  extends: Line,
  mixins: [reactiveProp],
  props: ['options'],
  mounted () {
    this.renderChart({
      labels: ['1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12', '13', '14', '1500'],
      datasets: [
        {
          label: 'Data One',
          backgroundColor: '#f87979',
          data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 90, 300, 1000, 1500]
        }
      ]
    }, {responsive: true, maintainAspectRatio: false})
  }
}


*/
