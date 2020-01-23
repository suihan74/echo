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
        <v-app>
          <v-app-bar app={true} color="primary" dark={true}>
            <Header propsData={this.$attrs} />
          </v-app-bar>
          <v-content>
            <PageComponent propsData={this.$attrs} />
          </v-content>
        </v-app>
      )
    }
  })
}
