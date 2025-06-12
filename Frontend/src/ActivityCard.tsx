import React from 'react';
import { ActivityListItem } from './types';

interface ActivityCardProps {
    activity: ActivityListItem;
    onClick: (id: number) => void;
}

const ActivityCard: React.FC<ActivityCardProps> = ({ activity, onClick }) => {
    return (
        <div className="activity-card" onClick={() => onClick(activity.id)}>
            <div className="activity-card-header">
                <h3 className="activity-title">{activity.titulo}</h3>
                <span className="activity-time">{activity.horario}</span>
            </div>
            <div className="activity-card-body">
                <p className="activity-instructor">
                    <strong>Profesor:</strong> {activity.profesor}
                </p>
            </div>
            <div className="activity-card-footer">
                <button className="btn-details">Ver detalles</button>
            </div>
        </div>
    );
};

export default ActivityCard;