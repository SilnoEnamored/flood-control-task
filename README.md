# Flood control TASK
Я сделла в двух вариантах, с использованием Redis и PostgresQL:  
1. Redis использовал так как он хорош в качестве быстрого хранилища, он предоставляет необходимую простоту и скорость реализации. Первый опыт взаимодействия с Redis, так что было интересно разобраться и попробовать.
2. PostgresQL больше подходит для долговременного хранения, более надежен. Уже был опыт использования.
   
Также добавил поддержку конфигураций с помощью viper (github.com/spf13/viper v1.18.2), так как уже был опыт использования этой бибилотеки и для данной задачи она отлично подходит.
