import Vue from 'vue'
import Header from '@/components/Header'

export default function simpleLayout (PageComponent) {
  return Vue.extend({
    inheritAttrs: false,
    components: {
      Header,
      PageComponent
    },
    render () {
      return (
        <div>
          <Header propsData={this.$attrs} />
          <PageComponent propsData={this.$attrs} />
        </div>
      )
    }
  })
}
