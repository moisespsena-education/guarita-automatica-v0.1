package main

import (
	"fmt"
	"time"
)

type Cancela struct {
	Aberta bool
}

func (c *Cancela) Abre() {
	fmt.Println("Abrindo cancela...")
	c.Aberta = true
	fmt.Println("Cancela aberta.")
}

func (c *Cancela) Fecha() {
	fmt.Println("Fechando cancela...")
	c.Aberta = false
	fmt.Println("Cancela fechada.")
}

type Sensor struct {
	qnt int
}

func (s *Sensor) VemCarro() bool {
	defer func() {
		s.qnt++
	}()
	if (s.qnt % 2) == 0 {
		return true
	}
	return false
}

type SensorPassagem struct {
	qnt int
}

func (s *SensorPassagem) Passou() bool {
	defer func() {
		s.qnt++
	}()
	if (s.qnt % 3) == 0 {
		return true
	}
	return false
}

type Porteiro struct {
	Cancela *Cancela
	Sensor Sensor
	SensorPassagem SensorPassagem
}

func (p *Porteiro) Trabalha() {
	for {
		now := time.Now()
		fmt.Println(now)
		if now.Hour() >= 7 && (now.Hour() <= 20 && now.Minute() < 54) {
			fmt.Println("Começa.")

			for !p.Sensor.VemCarro() {
				fmt.Println("Nao vem carro.")
				<-time.After(time.Second * 3)
			}

			fmt.Println("Vem carro.")

			p.Cancela.Abre()
			<-time.After(time.Second * 3)

			for !p.SensorPassagem.Passou() {
				fmt.Println("Nao passou.")
				<-time.After(time.Second * 3)
			}

			fmt.Println("Passou.")

			p.Cancela.Fecha()
			<-time.After(time.Second * 3)
		} else {
			fmt.Println("fora do horario de servico")
			<-time.After(time.Second * 3)
		}
	}
}

func main() {
	sensor := Sensor{}
	sensorPassagem := SensorPassagem{}
	cancela := &Cancela{}
	porteiro := Porteiro{cancela, sensor, sensorPassagem}
	porteiro.Trabalha()
}
