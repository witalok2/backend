import { http } from './config'

export default {
  create:(atividade)=>{
		return http.post('criar', atividade);
  },
    
	update:(atividade)=>{
		return http.put('atualizar/', atividade);
  },

  list:()=>{
		return http.get('listar')
  },
    
	apagar:(atividade)=>{
		return http.delete('deletar/' + atividade);
  }
}