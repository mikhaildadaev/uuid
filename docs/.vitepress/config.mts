import { defineConfig } from 'vitepress'
export default defineConfig({
  appearance: 'dark',
  base: '/uuid/',
  head: [
    ['link', { rel: 'stylesheet', href: '/uuid/styles.css' }],
    ['script', { src: '/uuid/scripts.js' }]
  ],
  lastUpdated: true,
  locales: {
    en: {
      description: 'A high-performance, zero-dependency platform for logs, metrics and traces.',
      label: 'English',
      lang: 'en',
      link: '/en/',
      title: 'UUID',
      themeConfig: {
        nav: [
          { 
            text: 'Home', 
            link: '/en/' 
          },
          { 
            text: 'Go', 
            link: '/en/go' 
          },
          { 
            text: 'Benchmarks', 
            link: '/en/benchmarks' 
          },
          { 
            text: 'API', 
            link: '/en/core_constructors' 
          },
        ],
        sidebar: [
          {
            items: [
              { 
                text: 'Go', 
                link: '/en/go' 
              },
              { 
                text: 'Benchmarks', 
                link: '/en/benchmarks' 
              },
              { 
                text: 'API', 
                collapsed: true,
                items: [
                  { 
                    text: 'Core', 
                    collapsed: true,
                    items: [
                      { 
                        text: 'Constructors', 
                        link: '/en/core_constructors' 
                      },
                      { 
                        text: 'Methods', 
                        link: '/en/core_methods'
                      } 
                    ] 
                  },
                  {
                    text: 'Marshal',
                    collapsed: true, 
                    items: [
                      { 
                        text: 'Methods', 
                        link: '/en/marshal_methods' 
                      }
                    ]
                  },
                  {
                    text: 'SQL', 
                    collapsed: true,
                    items: [
                      { 
                        text: 'Methods', 
                        link: '/en/sql_methods' 
                      }
                    ]
                  }
                ] 
              }
            ]
          }
        ],
        darkModeSwitchLabel: "Appearance",
        darkModeSwitchTitle: "Switch to dark theme",
        lightModeSwitchTitle: "Switch to light theme",
        sidebarMenuLabel: "Menu",
        returnToTopLabel: "Return to top",
        outline: {
          label: "On this page"
        },
        lastUpdated: {
          text: "Last Updated",
          formatOptions: {
            dateStyle: "short",
            timeStyle: "short"
          }
        },
        docFooter: {
          prev: "Previous page",
          next: "Next page"
        },
        footer: {
          message: 'Released under the Apache License 2.0',
          copyright: '© 2026 Mikhail Dadaev'
        }
      }
    },
    ru: {
      description: 'Производительная платформа без зависимостей для логов, метрик и трейсов.',
      label: 'Русский',
      lang: 'ru',
      link: '/ru/',
      title: 'UUID',
      themeConfig: {
        nav: [
          { 
            text: 'Главная', 
            link: '/ru/' 
          },
          { 
            text: 'Go', 
            link: '/ru/go' 
          },
          { 
            text: 'Бенчмарки', 
            link: '/ru/benchmarks' 
          },
          { 
            text: 'API', 
            link: '/ru/core_constructors' 
          },
        ],
        sidebar: [
          {
            items: [
              { 
                text: 'Go', 
                link: '/ru/go' 
              },
              { 
                text: 'Бенчмарки', 
                link: '/ru/benchmarks' 
              },
              { 
                text: 'API', 
                collapsed: true,
                items: [
                  { 
                    text: 'Ядро', 
                    collapsed: true,
                    items: [
                      { 
                        text: 'Конструкторы', 
                        link: '/ru/core_constructors' 
                      },
                      { 
                        text: 'Методы', 
                        link: '/ru/core_methods'
                      } 
                    ] 
                  },
                  { 
                    text: 'Сериализация', 
                    collapsed: true,
                    items: [
                      { 
                        text: 'Методы', 
                        link: '/ru/marshal_methods' 
                      }
                    ] 
                  },
                  { 
                    text: 'Интеграция с SQL', 
                    collapsed: true,
                    items: [
                      { 
                        text: 'Методы', 
                        link: '/ru/sql_methods' 
                      }
                    ] 
                  }
                ]
              }
            ]
          }
        ],
        darkModeSwitchLabel: "Внешний вид",
        darkModeSwitchTitle: "Переключиться на тёмную тему",
        lightModeSwitchTitle: "Переключиться на светлую тему",
        sidebarMenuLabel: "Меню",
        returnToTopLabel: "Вернуться наверх",
        outline: {
          label: "Содержание страницы"
        },
        lastUpdated: {
          text: "Последние изменения",
          formatOptions: {
            dateStyle: "short",
            timeStyle: "short"
          }
        },
        docFooter: {
          prev: "Предыдущая страница",
          next: "Следующая страница"
        },
        footer: {
          message: 'Под лицензией Apache 2.0',
          copyright: '© 2026 Михаил Дадаев'
        },
      }
    },
    zh: {
      description: '一个高性能、零依赖性的日志、度量和跟踪平台。',
      label: '简体中文',
      lang: 'zh',
      link: '/zh/',
      title: 'UUID',
      themeConfig: {
        nav: [
          { 
            text: '首页', 
            link: '/zh/' 
          },
          { 
            text: 'Go', 
            link: '/zh/go' 
          },
          { 
            text: '基准测试', 
            link: '/zh/benchmarks' 
          },
          { 
            text: 'API', 
            link: '/zh/core_constructors' 
          },
        ],
        sidebar: [
          {
            items: [
              { 
                text: 'Go', 
                link: '/zh/go' 
              },
              { 
                text: '基准', 
                link: '/zh/benchmarks' 
              },
              { 
                text: 'API', 
                collapsed: true,
                items: [
                  { 
                    text: '核心', 
                    collapsed: true,
                    items: [
                      { 
                        text: '构造函数', 
                        link: '/zh/core_constructors' 
                      },
                      { 
                        text: '方法', 
                        link: '/zh/core_methods'
                      }
                    ] 
                  },
                  { 
                    text: '序列化', 
                    collapsed: true,
                    items: [
                      { 
                        text: '方法', 
                        link: '/zh/marshal_methods' 
                      }
                    ] 
                  },
                  { 
                    text: 'SQL 集成', 
                    collapsed: true,
                    items: [
                      { 
                        text: '方法', 
                        link: '/zh/sql_methods' 
                      }
                    ] 
                  }
                ] 
              }
            ]
          }
        ],
        darkModeSwitchLabel: "深色模式",
        darkModeSwitchTitle: "切换至深色主题",
        lightModeSwitchTitle: "切换至浅色主题",
        sidebarMenuLabel: "目录",
        returnToTopLabel: "返回至顶部",
        outline: {
          label: "页面导航"
        },
        lastUpdated: {
          text: "最近更改",
          formatOptions: {
            dateStyle: "short",
            timeStyle: "short"
          }
        },
        docFooter: {
          prev: "上一页",
          next: "下一页"
        },
        footer: {
          message: '根据 Apache 2.0 许可证发布',
          copyright: '© 2026 Mikhail Dadaev'
        },
      }
    }
  },
  themeConfig: {
    search: {
      provider: 'local'
    },
    socialLinks: [
      { 
        icon: 'github', 
        link: 'https://github.com/mikhaildadaev/uuid' 
      }
    ],
  }
})