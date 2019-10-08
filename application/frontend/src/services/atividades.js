import { http } from './config'

export default {
  salvar:(atividade)=>{
		return http.post('criar', atividade);
  },
    
	atualizar:(atividade)=>{
		return http.put('atualizar/:id', atividade);
  },

  listar:()=>{
		return http.get('listar')
  },
    
	apagar:(atividade)=>{
		return http.delete('deletar/:id', { data: atividade })
  }
}