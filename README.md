# golang_pets_projects

## Вопросы на интервью:
- в чем разница между слайс и массивом? 
    [5]type, []type 
- Какая алгаритмическая сложность доступа по ключу к map? 
    Бакеты, Хэш-функции, Эвакуации данных
- Какие существуют каналы в Go? 
    Не буферизованный, буферизованный, закрытый, каналы для односторонней, не инициализированный канал. 
- Для чего нужен контекст в Go? 
    context.Background()
    context.TODO()
    context.WithCancel(parent Context)
    context.WithTimeout(parent Context, timeout time.Duration)
    context.WithDeadline(parent Context, deadline time.Time)
    context.WithValue(parent Context, key, value)
- Варианты механизмов для синхронизации работы горутины
    Каналы
    sync.Mutex
    sync.RWMutex
    sync.WaitGroup
    sync.Cond
    sync.Once
    atomic 
- Базы данных: Какие есть индексы?
    B-Tree (Балансированное дерево поиска)
    Hash Index
    GiST Index (Generalized Search Tree)
    SP-GiST Index (Space-partitioned Generalized Search Tree)
    GIN Index (Generalized Inverted Index)
    BRIN Index (Block Range INdexes)
    Partial Index
    Unique Index
- Базы данных: Что быстрее Index Scan и Sequential Scan? 
- Базы данных: Что такое транзакция? 
    Атомарность (Atomicity)
    Согласованность (Consistency)
    Изоляция (Isolation)
    Долговечность (Durability)
-  Базы данных: В чем разница между HAVING и WHERE?
-  Базы данных: Что такое тригеры? 
    До выполнения операции (BEFORE)
    После выполнения операции (AFTER)
    INSTEAD OF