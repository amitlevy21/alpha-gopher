import Users from './components/Users.vue'
import FileSystem from './components/FileSystem.vue'
import System from './components/System.vue'
import Help from './components/Help.vue'

export default [
    {path: '/users', component: Users},
    {path: '/filesystem', component: FileSystem},
    {path: '/system', component: System},
    {path: '/help', component: Help}
]