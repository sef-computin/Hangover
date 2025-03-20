package dbhandler

type EventDB struct{

}

type DBConfig struct{

}

func NewDBHandler(configurations DBConfig) *EventDB{
  return new(EventDB)
}
