package main

import "fmt"

//定义动物接口：所有动物都会新陈代谢，都会挂掉
type Animal interface {
	//新陈代谢：吃进来+排出去,shit就是翔啊~
	Eat(food string) (shit string)
	//GAME OVER
	Die()
}

//定义战士接口，会进攻和防守
type Fighter interface {
	//进攻并造成对手掉血
	Attack() (bloodLoss int)
	Defend()
}

//野兽接口，拥有动物的一切特征
//野兽接口，拥有战士的一切特征
type Beast interface {
	//野兽接口继承动物接口
	Animal
	//野兽接口继承斗士接口
	Fighter
	Run()
}

type Tiger struct {
	name  string
	food  string
	shit  string
	power int
}

/*实现野兽的全部方法才算是野兽*/
func (t *Tiger) Eat(food string) (shit string) {
	fmt.Printf("本王正在享用%s,并撇下%s\n", t.food, t.shit)
	return t.shit
}
func (t *Tiger) Die() {
	fmt.Printf("大猫%s也有狗带的一天啊，啊啊啊啊...\n", t.name)
}
func (t *Tiger) Attack() (bloodLoss int) {
	fmt.Printf("本王咬你，掉血%d毫升\n", t.power)
	return t.power
}
func (t *Tiger) Defend() {
	fmt.Println("躺在地上举高高，我不是在卖萌，而是在防守")
}
func (t *Tiger) Run() {
	fmt.Println("本王在奔跑")
}

func main() {

	tiger := &Tiger{"东北虎", "其它野兽", "虎翔", 1024}

	var animal Animal
	var fighter Fighter
	var beast Beast

	//老虎既是动物，又是斗士，又是野兽
	//用子类实现去给父类接口赋值
	animal = tiger
	fighter = tiger
	beast = tiger

	//调用父类接口方法
	animal.Eat("食物")
	animal.Die()
	fighter.Attack()
	fighter.Defend()
	beast.Run()
}
